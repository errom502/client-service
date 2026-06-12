import type { CSSProperties } from 'react'

export const s: Record<string, CSSProperties> = {
  page: {
    maxWidth: '480px',
    margin: '3rem auto',
    padding: '0 1rem',
    fontFamily: "'Inter', sans-serif",
    color: '#1f2937',
  },

  title: {
    fontSize: '1.75rem',
    fontWeight: 700,
    marginBottom: '0.5rem',
  },

  subtitle: {
    color: '#6b7280',
    marginBottom: '2rem',
  },

  cardGrid: {
    display: 'flex',
    flexDirection: 'column',
    gap: '1rem',
  },

  card: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    gap: '0.25rem',
    padding: '1.25rem',
    borderRadius: '0.75rem',
    border: '1px solid #dfe0e3ff',
    background: '#111827',
    cursor: 'pointer',
    textAlign: 'left',
    transition: 'box-shadow .15s',
    boxShadow: '0 1px 3px rgba(0,0,0,0.06)',
  },

  cardIcon: { fontSize: '1.5rem' },

  cardTitle: {
    fontWeight: 600,
    fontSize: '1rem',
    color: '#ffffff',
  },

  cardDesc: {
    fontSize: '0.875rem',
    color: '#6b7280',
  },

  form: {
    display: 'flex',
    flexDirection: 'column',
    gap: '0.75rem',
    marginBottom: '1rem',
  },

  input: {
    padding: '0.75rem 1rem',
    borderRadius: '0.5rem',
    border: '1px solid #d1d5db',
    fontSize: '1rem',
    outline: 'none',
  },

  btnPrimary: {
    padding: '0.75rem 1rem',
    borderRadius: '0.5rem',
    background: '#2563eb',
    color: '#fff',
    border: 'none',
    fontSize: '1rem',
    fontWeight: 600,
    cursor: 'pointer',
  },

  btnBack: {
    marginTop: '1.5rem',
    background: 'none',
    border: 'none',
    color: '#2563eb',
    cursor: 'pointer',
    fontSize: '0.95rem',
    padding: 0,
  },

  infoBox: {
    background: '#c6daeeff',
    border: '1px solid #e5e7eb',
    borderRadius: '0.5rem',
    padding: '1rem 1.25rem',
    marginBottom: '1.5rem',
    fontSize: '0.9rem',
  },

  infoTitle: {
    fontWeight: 600,
    marginBottom: '0.5rem',
    marginTop: 0,
  },

  infoList: {
    margin: 0,
    paddingLeft: '1.25rem',
    lineHeight: '1.8',
  },
}