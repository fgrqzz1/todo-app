import type { Task, CreateTaskInput } from '@/types/task'

/** Базовый URL API. В dev через Vite proxy запросы идут на бэкенд (см. vite.config proxy). */
const API_BASE = '/api'

async function request<T>(
  path: string,
  options: RequestInit = {}
): Promise<T> {
  const url = path.startsWith('http') ? path : `${API_BASE}${path}`
  const res = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
  })
  if (!res.ok) {
    const err = await res.json().catch(() => ({}))
    throw new Error((err as { error?: string }).error || res.statusText)
  }
  if (res.status === 204 || res.headers.get('content-length') === '0') {
    return undefined as T
  }
  return res.json() as Promise<T>
}

export const api = {
  getTasks: () => request<Task[]>('/tasks'),
  getTask: (id: number) => request<Task>(`/tasks/${id}`),
  createTask: (body: CreateTaskInput) =>
    request<Task>('/tasks', { method: 'POST', body: JSON.stringify(body) }),
  markDone: (id: number) =>
    request<{ message: string }>(`/tasks/${id}/done`, { method: 'PATCH' }),
  deleteTask: (id: number) =>
    request<{ message: string }>(`/tasks/${id}`, { method: 'DELETE' }),
}
