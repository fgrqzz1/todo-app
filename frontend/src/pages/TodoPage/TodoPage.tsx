import { useState, useEffect, useCallback } from 'react'
import { api } from '@/api/client'
import type { Task } from '@/types/task'
import { Layout } from '@/components/Layout/Layout'
import { Header } from '@/components/Header/Header'
import { TaskForm } from '@/components/TaskForm/TaskForm'
import { TaskList } from '@/components/TaskList/TaskList'
import styles from './TodoPage.module.css'

export function TodoPage() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [updatingId, setUpdatingId] = useState<number | null>(null)
  const [creating, setCreating] = useState(false)

  const fetchTasks = useCallback(async () => {
    setLoading(true)
    setError(null)
    try {
      const data = await api.getTasks()
      setTasks(data)
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Не удалось загрузить задачи')
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    fetchTasks()
  }, [fetchTasks])

  const handleCreate = async (title: string, description: string) => {
    setCreating(true)
    setError(null)
    try {
      const created = await api.createTask({ Title: title, Description: description || undefined })
      setTasks((prev) => [created, ...prev])
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Не удалось добавить задачу')
    } finally {
      setCreating(false)
    }
  }

  const handleToggleDone = async (id: number) => {
    setUpdatingId(id)
    setError(null)
    try {
      await api.markDone(id)
      setTasks((prev) =>
        prev.map((t) => (t.ID === id ? { ...t, Done: true } : t))
      )
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Не удалось обновить задачу')
    } finally {
      setUpdatingId(null)
    }
  }

  const handleDelete = async (id: number) => {
    setUpdatingId(id)
    setError(null)
    try {
      await api.deleteTask(id)
      setTasks((prev) => prev.filter((t) => t.ID !== id))
    } catch (e) {
      setError(e instanceof Error ? e.message : 'Не удалось удалить задачу')
    } finally {
      setUpdatingId(null)
    }
  }

  return (
    <Layout>
      <Header />
      {error && (
        <div className={styles.error} role="alert">
          {error}
        </div>
      )}
      <main className={styles.main}>
        <section className={styles.listSection} aria-labelledby="tasks-heading">
          <h2 id="tasks-heading" className={styles.listHeading}>Задачи</h2>
          {loading ? (
            <div className={styles.loading}>Загрузка…</div>
          ) : (
            <TaskList
              tasks={tasks}
              onToggleDone={handleToggleDone}
              onDelete={handleDelete}
              updatingId={updatingId}
            />
          )}
        </section>
        <aside className={styles.formSection} aria-label="Добавить задачу">
          <TaskForm onSubmit={handleCreate} isLoading={creating} />
        </aside>
      </main>
    </Layout>
  )
}
