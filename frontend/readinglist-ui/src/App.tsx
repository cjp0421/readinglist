import './App.css'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import BookDetail from './components/BookDetail'
import Books from './components/Books'
import AddBook from './components/AddBook'

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/books/:bookId" element={<BookDetail />} />
        <Route path="/" element={<Books />} />
        <Route path="/books/add" element={<AddBook />} />
      </Routes>
    </Router>
  )
}

export default App
