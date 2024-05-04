import { useEffect, useState } from 'react';
import { getBookById } from '../services/bookService';
import { useParams } from 'react-router-dom';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import { Paper, Table, TableBody, TableHead, TableRow } from '@mui/material';

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
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        setIsLoading(true)
        console.log('Fetching book with ID: ', bookId)
        getBookById(Number(bookId))
            .then(data => {
                console.log('Book data received: ', data)
                setBook(data.book)
                setIsLoading(false)
                console.log(book)
            })
            .catch(error => {
                console.error('Error fetching book:', error);
                setIsLoading(false)
            }
            )

    }, [bookId])

    if (isLoading) {
        return <p>Loading</p>
    }


    return (

        <TableContainer component={Paper}>
            <Table aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell className="columnHeader" aria-label="ID">ID</TableCell>
                        <TableCell className="columnHeader" aria-label="Title">Title </TableCell>
                        <TableCell className="columnHeader" aria-label="Author">Author </TableCell>
                        <TableCell className="columnHeader" aria-label="Published">Published </TableCell>
                        <TableCell className="columnHeader" aria-label="Pages">Pages </TableCell>
                        <TableCell className="">Genres</TableCell>
                        <TableCell className="columnHeader" aria-label="Rating">Rating </TableCell>
                        <TableCell className="">ISBN</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>

                    <TableRow key={book.id}>
                        <TableCell component="th" scope="row">{book.id}</TableCell>
                        <TableCell>{book.title}</TableCell>
                        <TableCell>{book.author}</TableCell>
                        <TableCell>{book.published}</TableCell>
                        <TableCell>{book.pages}</TableCell>
                        <TableCell>{book.genres.join(', ')}</TableCell>
                        <TableCell>{book.rating.toFixed(1)}</TableCell>
                        <TableCell>{book.isbn}</TableCell>
                    </TableRow>

                </TableBody>
            </Table>
        </TableContainer>
    )
}

export default BookDetail;