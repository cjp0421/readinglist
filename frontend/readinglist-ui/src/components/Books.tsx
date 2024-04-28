import React, { useEffect, useState } from "react";
import { getBooks } from "../services/bookService";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';

interface Book {
    id: number;
    title: string;
    author: string;
    published: string;
    pages: number;
    genres: string[];
    rating: number;
}

export const Books: React.FC = () => {
    const [books, setBooks] = useState<Book[]>([]);
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        getBooks().then(response => {
            console.log("API Response:", response)
            if (Array.isArray(response.books)) {
                setBooks(response.books);
            } else {
                console.error('Data received is not an array:', response.books)
            }
        }).catch(error => setError('Failed to fetch books: ' + error))
    }, []);

    if (error) return <div>Error: {error}</div>
    if (!books.length) return <div>Loading books...</div>

    return (
        <div>
            <h1>Book List</h1>
            <TableContainer component={Paper}>
                <Table aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>ID</TableCell>
                            <TableCell>Title</TableCell>
                            <TableCell>Author</TableCell>
                            <TableCell>Pubished</TableCell>
                            <TableCell>Pages</TableCell>
                            <TableCell>Genres</TableCell>
                            <TableCell>Rating</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {books.map((book) => (
                            <TableRow key={book.id}>
                                <TableCell component="th" scope="row">{book.id}</TableCell>
                                <TableCell>{book.title}</TableCell>
                                <TableCell>{book.author}</TableCell>
                                <TableCell>{book.published}</TableCell>
                                <TableCell>{book.pages}</TableCell>
                                <TableCell>{book.genres.join(', ')}</TableCell>
                                <TableCell>{book.rating.toFixed(1)}</TableCell>
                            </TableRow>
                        )
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        </div>
    )
};

export default Books;