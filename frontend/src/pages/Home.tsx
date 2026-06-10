import { useNavigate } from 'react-router-dom'
import { s } from '../styles'

export function Home() {
  const nav = useNavigate()
  return (
    <div style={s.page}>
      <h1 style={s.title}>Система верификации почты</h1>
      <p style={s.subtitle}>Выберите действие</p>
      <div style={s.cardGrid}>
        <button style={s.card} onClick={() => nav('/status')}>
          <span style={s.cardIcon}>🔍</span>
          <span style={s.cardTitle}>Проверить статус</span>
          <span style={s.cardDesc}>Узнать, верифицирована ли почта</span>
        </button>
        <button style={s.card} onClick={() => nav('/create')}>
          <span style={s.cardIcon}>✉️</span>
          <span style={s.cardTitle}>Верифицировать почту</span>
          <span style={s.cardDesc}>Создать новую верификацию</span>
        </button>
        <button style={s.card} onClick={() => nav('/retry')}>
          <span style={s.cardIcon}>🔄</span>
          <span style={s.cardTitle}>Повторная верификация</span>
          <span style={s.cardDesc}>Отправить письмо ещё раз</span>
        </button>
      </div>
    </div>
  )
}
