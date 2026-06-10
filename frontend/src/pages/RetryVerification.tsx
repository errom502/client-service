import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { api } from '../api/verification'
import { Notification } from '../components/Notification'
import { s } from '../styles'

export function RetryVerification() {
  const nav = useNavigate()
  const [email, setEmail] = useState('')
  const [loading, setLoading] = useState(false)
  const [success, setSuccess] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)

  const handle = async () => {
    if (!email) return
    setLoading(true)
    setSuccess(null)
    setError(null)
    try {
      const resp = await api.retry(email, email)
      setSuccess(resp.message)
      setEmail('')
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={s.page}>
      <h2 style={s.title}>Повторная верификация</h2>

      {/* Описание логики для пользователя */}
      <div style={s.infoBox}>
        <p style={s.infoTitle}>Как это работает:</p>
        <ul style={s.infoList}>
          <li>Если верификация для вашей почты <b>не найдена</b> — создаётся новая</li>
          <li>Если верификация <b>подтверждена</b> — создаётся новая верификация</li>
          <li>Если верификация <b>истекла или не удалась</b> — отправляется новое письмо для существующей</li>
          <li>Если верификация <b>активна (ожидает подтверждения)</b> — повтор недоступен, проверьте почту</li>
        </ul>
      </div>

      <div style={s.form}>
        <input
          style={s.input}
          type="email"
          placeholder="Введите email"
          value={email}
          onChange={e => setEmail(e.target.value)}
          onKeyDown={e => e.key === 'Enter' && handle()}
        />
        <button style={s.btnPrimary} onClick={handle} disabled={loading || !email}>
          {loading ? 'Обрабатываю...' : 'Повторить верификацию'}
        </button>
      </div>

      {success && <Notification type="success" message={success} />}
      {error   && <Notification type="error"   message={error}   />}

      <button style={s.btnBack} onClick={() => nav('/')}>← Назад</button>
    </div>
  )
}
