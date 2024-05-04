import './App.css'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import BookDetail from './components/BookDetail'
import Books from './components/Books'

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/books/:bookId" element={<BookDetail />} />
        <Route path="/" element={<Books />} />
      </Routes>
    </Router>
  )
}

export default App
