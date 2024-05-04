import { useForm } from 'react-hook-form';
import { createBook } from '../services/bookService';
import { useNavigate } from 'react-router-dom';

const AddBook = () => {
    const { register, handleSubmit, formState: { errors } } = useForm();
    const navigate = useNavigate();

    const onSubmit = async (data) => {
        console.log(data)

        const formattedData = {
            title: data.title,
            author: data.author,
            published: parseInt(data.published, 10),
            pages: parseInt(data.pages, 10),
            genres: data.genres.split(',').map(genre => genre.trim()),
            rating: parseFloat(data.rating),
            isbn: data.isbn
        }

        console.log(formattedData)

        try {
            const newBook = await createBook(formattedData);
            console.log('Book created:', newBook)
            navigate('/')
        } catch (error) {
            console.error('Failed to create book:', error)
        }
    }

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <label htmlFor='title'>Title</label>
            <input id='title' {...register('title',
                { required: true }
            )} />
            {errors.title && <span>This field is required.</span>}

            <label htmlFor='author'>Author</label>
            <input id='author' {...register('author',
                { required: true }
            )} />
            {errors.author && <span>This field is required.</span>}

            <label htmlFor='published'>Published</label>
            <input id='published' {...register('published',
                { required: true }
            )} />
            {errors.published && <span>This field is required.</span>}

            <label htmlFor='pages'>Number of Pages</label>
            <input id='pages' {...register('pages',
                { required: true }
            )} />
            {errors.pages && <span>This field is required.</span>}

            <label htmlFor='Genres'>Genres</label>
            <input id='genres' {...register('genres',
                { required: true }
            )} />
            {errors.genres && <span>This field is required.</span>}

            <label htmlFor='rating'>Rating</label>
            <input id='rating' {...register('rating',
                { required: true }
            )} />
            {errors.rating && <span>This field is required.</span>}

            <label htmlFor='isbn'>ISBN</label>
            <input id='isbn' {...register('isbn')} />

            <button type='submit'>Add Book</button>
        </form>
    )
}

export default AddBook;