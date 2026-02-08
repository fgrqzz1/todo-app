import { useState, useEffect, useCallback } from 'react'
import { Layout } from '@/components/Layout/Layout'
import { Header } from '@/components/Header/Header'
import { TaskList } from '@/components/TaskList/TaskList'
import { TaskForm } from '@/components/TaskForm/TaskForm'
import { api } from '@/api/client'
import type { Task } from '@/types/task'
import styles from './TodoPage.module.css'

export function TodoPage() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  const [loadError, setLoadError] = useState<string | null>(null)
  const [updatingId, setUpdatingId] = useState<number | null>(null)
  const [formLoading, setFormLoading] = useState(false)
  const [formError, setFormError] = useState<string | null>(null)

  const loadTasks = useCallback(async () => {
    try {
      setLoading(true)
      setLoadError(null)
      const data = await api.getTasks()
      setTasks(data)
    } catch (e) {
      const msg = e instanceof Error ? e.message : 'Ошибка сети'
      setLoadError(msg)
      console.error('Failed to load tasks:', e)
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => {
    loadTasks()
  }, [loadTasks])

  const handleCreate = async (title: string, description: string) => {
    try {
      setFormLoading(true)
      setFormError(null)
      const created = await api.createTask({ Title: title, Description: description || undefined })
      setTasks((prev) => [created, ...prev])
    } catch (e) {
      const msg = e instanceof Error ? e.message : 'Не удалось добавить задачу'
      setFormError(msg)
      console.error('Failed to create task:', e)
    } finally {
      setFormLoading(false)
    }
  }

  const handleToggleDone = async (id: number) => {
    try {
      setUpdatingId(id)
      await api.markDone(id)
      setTasks((prev) =>
        prev.map((t) => (t.ID === id ? { ...t, Done: true } : t))
      )
    } catch (e) {
      console.error('Failed to mark done:', e)
    } finally {
      setUpdatingId(null)
    }
  }

  const handleDelete = async (id: number) => {
    try {
      setUpdatingId(id)
      await api.deleteTask(id)
      setTasks((prev) => prev.filter((t) => t.ID !== id))
    } catch (e) {
      console.error('Failed to delete task:', e)
    } finally {
      setUpdatingId(null)
    }
  }

  return (
    <Layout>
      <Header />
      <main className={styles.main}>
        <section className={styles.section} aria-labelledby="tasks-heading">
          <h2 id="tasks-heading" className={styles.heading}>
            Задачи
          </h2>
          {loading ? (
            <div className={styles.loading}>Загрузка…</div>
          ) : loadError ? (
            <div className={styles.error}>
              <p className={styles.errorText}>{loadError}</p>
              <p className={styles.errorHint}>Убедитесь, что бэкенд запущен: <code>make backend</code> или <code>make dev</code></p>
              <button type="button" className={styles.retry} onClick={loadTasks}>
                Повторить
              </button>
            </div>
          ) : (
            <TaskList
              tasks={tasks}
              onToggleDone={handleToggleDone}
              onDelete={handleDelete}
              updatingId={updatingId}
            />
          )}
        </section>
        <aside className={styles.aside} aria-labelledby="new-task-heading">
          <h2 id="new-task-heading" className={styles.heading}>
            Новая задача
          </h2>
          <div className={styles.formCard}>
            {formError && (
              <p className={styles.formError} role="alert">
                {formError}
              </p>
            )}
            <TaskForm onSubmit={handleCreate} isLoading={formLoading} />
          </div>
        </aside>
      </main>
    </Layout>
  )
}
