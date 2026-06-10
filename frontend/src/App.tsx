import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { CheckStatus } from './pages/CheckStatus'
import { CreateVerification } from './pages/CreateVerification'
import { Home } from './pages/Home'
import { RetryVerification } from './pages/RetryVerification'

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/"       element={<Home />}               />
        <Route path="/status" element={<CheckStatus />}        />
        <Route path="/create" element={<CreateVerification />} />
        <Route path="/retry"  element={<RetryVerification />}  />
      </Routes>
    </BrowserRouter>
  )
}