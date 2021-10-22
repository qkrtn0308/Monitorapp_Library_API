import react, {useState} from 'react'
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar'
import BookPageIndex from '../components/BookPageIndex';

const BookPage = () => {
    const [isOpen, setIsOpen] = useState(false);

    const toggle = () => {
        setIsOpen(!isOpen);
    };

    return (
        <>
            <Sidebar isOpen={isOpen} toggle={toggle} />
            <Navbar toggle={toggle} />
            <BookPageIndex/>
        </>
    );
};

export default BookPage;
