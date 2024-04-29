import React, { useMemo, useEffect, useState } from "react";
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

    const filteredData = useMemo(() => {
        if (searchTerm === "") {
            return books
        }
        return books
            .map((row) => {
                if (row.id.toString().includes(String(searchTerm).toLowerCase()) ||
                    row.title.includes(String(searchTerm).toLowerCase()) ||
                    row.author.includes(String(searchTerm).toLowerCase()) /*||
                    row.genres.toString().includes(String(searchTerm).toLowerCase())*/ ||
                    row.pages.toString().includes(String(searchTerm).toLowerCase()) ||
                    row.rating.toString().includes(String(searchTerm).toLowerCase())
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
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedData.map((row) => (
                            <TableRow key={row.id}>
                                <TableCell component="th" scope="row">{row.id}</TableCell>
                                <TableCell>{row.title}</TableCell>
                                <TableCell>{row.author}</TableCell>
                                <TableCell>{row.published}</TableCell>
                                <TableCell>{row.pages}</TableCell>
                                <TableCell>{row.genres.join(', ')}</TableCell>
                                <TableCell>{row.rating.toFixed(1)}</TableCell>
                                <TableCell>{row.isbn}</TableCell>
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