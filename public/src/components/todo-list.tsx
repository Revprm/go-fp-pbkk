/* eslint-disable @typescript-eslint/no-unused-vars */
"use client";

import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { taskSchema } from "@/lib/schemas";
import { Task } from "@/lib/definitions";
import { api } from "@/lib/api";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Trash2, Edit2 } from "lucide-react";
import type { z } from "zod";
import { useAuth } from "@/hooks/use-auth";

type TaskInputs = z.infer<typeof taskSchema>;

export function TodoList() {
  const router = useRouter();
  const { token, logout } = useAuth();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [editingTask, setEditingTask] = useState<Task | null>(null);

  const form = useForm<TaskInputs>({
    resolver: zodResolver(taskSchema),
    defaultValues: {
      title: "",
      description: "",
      status: "Open",
      due_date: new Date().toISOString().split("T")[0],
    },
  });

  const fetchTasks = async () => {
    if (!token) return;
    setIsLoading(true);
    try {
      // console.log(token);
      const response = await api.getTasks(token);
      if (response.status) {
        setTasks(response.data);
      }
    } catch (err) {
      setError("Failed to fetch tasks");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, [token]);

  const onSubmit = async (data: TaskInputs) => {
    if (!token) return;
    setIsLoading(true);
    try {
      if (editingTask) {
        await api.updateTask(token, editingTask.id, data);
      } else {
        await api.createTask(token, data);
      }
      form.reset();
      setEditingTask(null);
      fetchTasks();
    } catch (err) {
      setError(editingTask ? "Failed to update task" : "Failed to create task");
    } finally {
      setIsLoading(false);
    }
  };

  const handleDeleteTask = async (id: string) => {
    if (!token) return;
    try {
      await api.deleteTask(token, id);
      fetchTasks();
    } catch (err) {
      setError("Failed to delete task");
    }
  };

  const handleEditTask = (task: Task) => {
    setEditingTask(task);
    form.reset({
      title: task.title,
      description: task.description,
      status: task.status,
      due_date: new Date(task.due_date).toISOString().split("T")[0],
    });
  };

  const handleLogout = () => {
    logout();
    router.push("/login");
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">Task Manager</h1>
        <Button variant="outline" onClick={handleLogout}>
          Logout
        </Button>
      </div>

      {error && (
        <Alert variant="destructive">
          <AlertDescription>{error}</AlertDescription>
        </Alert>
      )}

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Title</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Description</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="status"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Status</FormLabel>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue placeholder="Select status" />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectItem value="Open">Open</SelectItem>
                    <SelectItem value="In Progress">In Progress</SelectItem>
                    <SelectItem value="Completed">Completed</SelectItem>
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            control={form.control}
            name="due_date"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Due Date</FormLabel>
                <FormControl>
                  <Input
                    type="date"
                    {...field}
                    value={new Date(field.value).toISOString().split("T")[0]}
                    onChange={(e) =>
                      field.onChange(new Date(e.target.value).toISOString())
                    }
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          <Button type="submit" disabled={isLoading}>
            {isLoading
              ? "Loading..."
              : editingTask
              ? "Update Task"
              : "Create Task"}
          </Button>
        </form>
      </Form>

      <div className="grid gap-4">
        {tasks.map((task) => (
          <div
            key={task.id}
            className="p-4 border rounded-lg shadow-sm hover:shadow-md transition-shadow"
          >
            <div className="flex justify-between items-start">
              <div className="space-y-1">
                <h3 className="font-semibold">{task.title}</h3>
                <p className="text-sm text-gray-600">{task.description}</p>
                <div className="flex gap-2 text-sm">
                  <span className="px-2 py-1 rounded-full bg-gray-100">
                    {task.status}
                  </span>
                  <span className="text-gray-500">
                    Due: {new Date(task.due_date).toLocaleDateString()}
                  </span>
                </div>
              </div>
              <div className="flex gap-2">
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={() => handleEditTask(task)}
                >
                  <Edit2 className="h-4 w-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon"
                  onClick={() => handleDeleteTask(task.id)}
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
