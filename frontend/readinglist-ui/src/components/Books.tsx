import React, { useMemo, useEffect, useState } from "react";
import { getBooks, deleteBook } from "../services/bookService";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { IconButton, Link } from "@mui/material";
import { Link as RouterLink } from 'react-router-dom';


interface Book {
    id: number;
    title: string;
    author: string;
    published: string;
    pages: number;
    genres: string[];
    rating: number;
    isbn: string;
}

export const Books: React.FC = () => {
    const [books, setBooks] = useState<Book[]>([]);
    const [error, setError] = useState<string | null>(null)
    const [searchTerm, setSearchTerm] = useState<string | number>("")
    const [sortColumn, setSortColumn] = useState<string>("id")
    const [sortOrder, setSortOrder] = useState<string>("asc")

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

    const handleHeaderClick = (column: string) => {
        if (column === sortColumn) {
            setSortOrder((prev) => (prev === "asc" ? "desc" : "asc"))
        } else {
            setSortColumn(column)
            setSortOrder("asc")
        }
    }

    const handleDelete = async (id) => {
        try {
            await deleteBook(id)
            setBooks(books.filter(book => book.id !== id))
            console.log('Book deleted successfully')
        } catch (error) {
            console.error('Error deleting book:', error)
        }
    }

    const filteredData = useMemo(() => {
        const lowerCaseSearchTerm = String(searchTerm).toLowerCase();
        if (searchTerm === "") {
            return books
        }
        return books
            .map((row) => {
                if (
                    row.id.toString().includes(lowerCaseSearchTerm) ||
                    row.title.toLowerCase().includes(lowerCaseSearchTerm) ||
                    row.author.toLowerCase().includes(lowerCaseSearchTerm) ||
                    row.genres.some(genre => genre.toLowerCase().includes(lowerCaseSearchTerm)) ||
                    row.pages.toString().includes(lowerCaseSearchTerm) ||
                    row.rating.toString().includes(lowerCaseSearchTerm) ||
                    row.isbn.toString().includes(lowerCaseSearchTerm)
                ) {
                    return row
                }
                return null;
            }).filter(Boolean)
    }, [searchTerm, books])


    const sortedData = useMemo(() => {
        const sorted = [...filteredData];

        return sorted.sort((a, b) => {
            const aValue = a[sortColumn]
            const bValue = b[sortColumn]

            if (sortOrder === "asc") {
                return aValue > bValue ? 1 : -1;
            } else {
                return aValue < bValue ? 1 : -1;
            }
        })
    }, [filteredData, sortColumn, sortOrder])

    if (error) return <div>Error: {error}</div>
    if (!books.length) return <div>Loading books...</div>

    return (
        <div>
            <header>
                <h1>Book List</h1>
                <input type="text" value={searchTerm} onChange={(e) => setSearchTerm(e.target.value)} placeholder="Search Books" />
            </header>
            <TableContainer component={Paper}>
                <Table aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("id")} aria-label="ID">ID {sortColumn === "id" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("title")} aria-label="Title">Title {sortColumn === "title" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("author")} aria-label="Author">Author {sortColumn === "author" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("published")} aria-label="Published">Published {sortColumn === "published" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("pages")} aria-label="Pages">Pages {sortColumn === "pages" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="">Genres</TableCell>
                            <TableCell className="columnHeader" onClick={() => handleHeaderClick("rating")} aria-label="Rating">Rating {sortColumn === "rating" && (sortOrder === "asc" ? "↑" : "↓")}</TableCell>
                            <TableCell className="">ISBN</TableCell>
                            <TableCell className="">Action</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedData.map((row) => (
                            <TableRow key={row?.id}>
                                <TableCell component="th" scope="row">{row?.id}</TableCell>
                                <TableCell>
                                    <Link component={RouterLink} to={`/books/${row?.id}`}>
                                        {row?.title}
                                    </Link>
                                </TableCell>
                                <TableCell>{row?.author}</TableCell>
                                <TableCell>{row?.published}</TableCell>
                                <TableCell>{row?.pages}</TableCell>
                                <TableCell>{row?.genres.join(', ')}</TableCell>
                                <TableCell>{row?.rating.toFixed(1)}</TableCell>
                                <TableCell>{row?.isbn}</TableCell>
                                <TableCell>
                                    <IconButton onClick={() => handleDelete(row?.id)} aria-label="delete">
                                        <FontAwesomeIcon icon={faTrash} />
                                    </IconButton>
                                </TableCell>
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