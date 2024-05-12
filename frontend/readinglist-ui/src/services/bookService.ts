/* eslint-disable no-useless-catch */
import axios from 'axios'

const API_URL = 'http://localhost:3001/v1/books';

export const getBooks = async () => {
    try {
        const response = await axios.get(API_URL);
        return response.data;
    } catch (error) {
        throw error;
    }
};

export const getBookById = async (id: number) => {
    try {
        const response = await axios.get(`${API_URL}/${id}`)
        console.log("API:", response.data.book)
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch book details');
    }
}

export const createBook = async (bookData) => {
    try {
        const response = await axios.post(API_URL, bookData);
        return response.data;
    } catch (error) {
        //todo: fix up this error handlingto be spiffier
        throw error;
    }
}

export const updateBookById = async (id: number, bookData) => {
    try {
        const response = await axios.put(`${API_URL}/${id}`, bookData)
        return response.data;
    } catch (error) {
        console.error('Failed to update book:', error);
        throw new Error('Failed to update book details')
    }
}

export const deleteBook = async (id: number) => {
    try {
        const response = await axios.delete(`${API_URL}/${id}`)
        return response.data;
    } catch (error) {
        console.error('Failed to delete book:', error)
        throw error;
    }
}