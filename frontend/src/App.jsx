import './App.css'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import AppRoutes from './routes/AppRoutes'
import { Suspense } from 'react'
import AppSkeleton from './modules/skeleton/AppSkeleton'

function App() {
  return (
    <>
      <BrowserRouter>
      <Suspense fallback={<><AppSkeleton></AppSkeleton><AppSkeleton></AppSkeleton></>}>
        <Routes>
          <Route path="/*" element={<AppRoutes/>}></Route>
        </Routes>
      </Suspense>
      </BrowserRouter>
    </>
  )
}

export default App
