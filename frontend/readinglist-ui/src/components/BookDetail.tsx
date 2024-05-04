import { useEffect, useState } from 'react';
import { getBookById } from '../services/bookService';
import { useParams } from 'react-router-dom';

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
            .then(book => {
                console.log('Book data received: ', book)
                setBook(book)
            })
            .catch(error => {
                console.error('Error fetching book:', error);
                setIsLoading(false)
            }
            )
        setIsLoading(false)
    }, [bookId])

    if (isLoading) {
        return <p>Loading</p>
    }


    return (

        <div>
            {console.log("Rendering:", book)}
            <div>
                <h1>{book?.title}</h1>
                <p>{book?.id}</p>
            </div>
        </div>
    )
}

export default BookDetail;