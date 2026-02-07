import { useState, FormEvent } from 'react'
import styles from './TaskForm.module.css'

interface TaskFormProps {
  onSubmit: (title: string, description: string) => void
  isLoading?: boolean
}

export function TaskForm({ onSubmit, isLoading }: TaskFormProps) {
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()
    const t = title.trim()
    if (!t) return
    onSubmit(t, description.trim())
    setTitle('')
    setDescription('')
  }

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <div className={styles.inputs}>
        <input
          type="text"
          className={styles.input}
          placeholder="Новая задача..."
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          disabled={isLoading}
          maxLength={200}
          autoComplete="off"
        />
        <input
          type="text"
          className={styles.inputDescription}
          placeholder="Описание (необязательно)"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          disabled={isLoading}
          maxLength={500}
          autoComplete="off"
        />
      </div>
      <button
        type="submit"
        className={styles.submit}
        disabled={isLoading || !title.trim()}
      >
        {isLoading ? 'Добавляю…' : 'Добавить'}
      </button>
      <p className={styles.credit}>
        Разработчик Grachev Yaroslav
        <br />
        tg:{' '}
        <a
          href="https://t.me/grachev_yk"
          target="_blank"
          rel="noopener noreferrer"
          className={styles.creditLink}
        >
          @grachev_yk
        </a>
      </p>
    </form>
  )
}
