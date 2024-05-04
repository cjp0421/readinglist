/* eslint-disable no-useless-catch */
import axios from 'axios'

const API_URL = 'http://localhost:4000/v1/books';

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
        console.log("API:", response.data.book.author)
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch book details');
    }
}