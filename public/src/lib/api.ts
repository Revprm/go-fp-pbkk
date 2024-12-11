import { ApiResponse, User, Task } from './definitions';
import { loginSchema, registerSchema, taskSchema } from './schemas';
import { z } from 'zod';

const API_BASE_URL = 'http://127.0.0.1:8888/api';

export class ApiError extends Error {
  constructor(public status: number, message: string) {
    super(message);
    this.name = 'ApiError';
  }
}

export const api = {
  async handleResponse<T>(response: Response): Promise<ApiResponse<T>> {
    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new ApiError(
        response.status,
        error.message || 'An error occurred'
      );
    }
    return response.json();
  },

  async login(data: z.infer<typeof loginSchema>) {
    const response = await fetch(`${API_BASE_URL}/user/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return this.handleResponse<{ token: string; role: string }>(response);
  },

  async register(data: z.infer<typeof registerSchema>) {
    const response = await fetch(`${API_BASE_URL}/user`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    return this.handleResponse<User>(response);
  },

  async getMe(token: string) {
    const response = await fetch(`${API_BASE_URL}/user/me`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return this.handleResponse<User>(response);
  },

  async getTasks(token: string) {
    const response = await fetch(`${API_BASE_URL}/task`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    return this.handleResponse<Task[]>(response);
  },

  async createTask(token: string, data: z.infer<typeof taskSchema>) {
    const response = await fetch(`${API_BASE_URL}/task`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    return this.handleResponse<Task>(response);
  },

  async updateTask(
    token: string,
    id: string,
    data: Partial<z.infer<typeof taskSchema>>
  ) {
    const response = await fetch(`${API_BASE_URL}/task/${id}`, {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    return this.handleResponse<Task>(response);
  },

  async deleteTask(token: string, id: string) {
    const response = await fetch(`${API_BASE_URL}/task/${id}`, {
      method: 'DELETE',
      headers: { Authorization: `Bearer ${token}` },
    });
    return this.handleResponse<void>(response);
  },
};