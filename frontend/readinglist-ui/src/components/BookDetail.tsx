import { useEffect, useState } from 'react';
import { getBookById, updateBookById } from '../services/bookService';
import { useParams } from 'react-router-dom';
import { TableCell, TableContainer, Paper, Table, TableBody, TableHead, TableRow, TextField } from '@mui/material';

interface Book {
    id: number;
    title: string;
    author: string;
    published: number;
    pages: number;
    genres: string[];
    rating: number;
    isbn: string;
}

const initialBookState: Book = {
    id: 0,
    title: 'Loading',
    author: 'Loading',
    published: 0,
    pages: 0,
    genres: [],
    rating: 0,
    isbn: 'Loading'
}

const BookDetail = () => {
    const { bookId } = useParams<{ bookId: string }>();
    const [book, setBook] = useState<Book>(initialBookState);
    const [editState, setEditState] = useState<Record<string, boolean>>({})

    useEffect(() => {
        console.log('Fetching book with ID: ', bookId)
        getBookById(Number(bookId))
            .then(data => {
                console.log('Book data received: ', data)
                setBook(data.book)
                console.log(book)
            })
            .catch(error => {
                console.error('Error fetching book:', error);
            }
            )

    }, [bookId])

    const toggleEdit = (field: string) => {
        setEditState(prev => ({ ...prev, [field]: !prev[field] }))
    }

    const handleChange = (field: string, value: string | number) => {
        if (book) {
            setBook(prev => prev ? { ...prev, [field]: value } : null)
        }
    }

    const handleBlur = (field: string) => {
        toggleEdit(field);
        if (book) {
            updateBookById(book.id, { [field]: book[field] }).catch(error => {
                console.error(`Failed to update ${field}:`, error);
            });
        }
    }

    if (!book) {
        return <p>Loading</p>
    }


    return (

        <TableContainer component={Paper}>
            <Table aria-label="simple table">
                <TableHead>
                    <TableRow>
                        {Object.keys(book).map(key => (
                            <TableCell key={key}>{(key != "id" && key != "isbn") ? key[0].toUpperCase() + key.substring(1) : key.toUpperCase()}</TableCell>
                        ))}
                    </TableRow>
                </TableHead>
                <TableBody>
                    <TableRow>

                        {Object.entries(book).map(([key, value]) =>
                        (
                            <TableCell key={key} onClick={() => toggleEdit(key)}>
                                {
                                    editState[key] ? (
                                        <TextField
                                            value={value}
                                            onChange={(e) => handleChange(key, e.target.value)}
                                            onBlur={() => handleBlur(key)}
                                            autoFocus
                                            fullWidth
                                        />


                                    ) : (
                                        key === 'genres' ? value.join(', ') : value
                                    )}


                            </TableCell>
                        ))}
                    </TableRow>
                </TableBody>
            </Table>
        </TableContainer >
    )
}

export default BookDetail;