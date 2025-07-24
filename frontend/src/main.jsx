import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import { Suspense } from 'react'
import Header from './modules/Header'

createRoot(document.getElementById('root')).render(
    <StrictMode>
      <Header />
      <Suspense fallback={<></>}>
        <App />
      </Suspense>
    </StrictMode>
)
