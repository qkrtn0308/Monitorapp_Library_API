import react, {useState} from 'react'
import Sidebar from '../components/Sidebar';
import Navbar from '../components/Navbar'
import Main from '../components/Test';
const AdminPage  = () => {

    const [isOpen, setIsOpen] = useState(false);

    const toggle = () => {
        setIsOpen(!isOpen);
    };

    return (
        <>
            <Main/>
        </>
    )
}

export default AdminPage;

