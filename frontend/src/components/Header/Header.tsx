import styles from './Header.module.css'

export function Header() {
  return (
    <header className={styles.header}>
      <div className={styles.logo}>
        <span className={styles.icon}>◇</span>
        <h1 className={styles.title}>Todo</h1>
      </div>
      <p className={styles.subtitle}>Список задач</p>
    </header>
  )
}
