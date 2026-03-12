import Button from 'react-bootstrap/Button'
import Container from 'react-bootstrap/Container'
import Nav from 'react-bootstrap/Nav'
import Navbar from 'react-bootstrap/Navbar'
import { NavLink, useNavigate } from 'react-router-dom'
import logo from '../../assets/MagicStreamLogo.png'
import useAuth from '../../hooks/useAuth'

const Header = ({ handleLogout }) => {
	// const Header = () => {
	const navigate = useNavigate()
	const { auth } = useAuth()
	// const [auth, _] = useState(false)

	return (
		<Navbar
			bg='light'
			data-bs-theme='light'
			expand='lg'
			// fixed='top'
			sticky='top'
			className='shadow-sm'
		>
			<Container>
				<Navbar.Brand>
					<img
						alt=''
						src={logo}
						width='30'
						height='30'
						className='d-inline-block align-top me-2'
					/>
					Magic Stream
				</Navbar.Brand>

				<Navbar.Toggle aria-controls='main-navbar-nav' />
				<Navbar.Collapse>
					<Nav className='me-auto'>
						<Nav.Link as={NavLink} to='/'>
							Home
						</Nav.Link>
						<Nav.Link as={NavLink} to='/recommended'>
							Recommended
						</Nav.Link>
					</Nav>

					<Nav className='ms-auto align-items-center'>
						{auth ? (
							<>
								<span className='me-3'>
									Hello, <strong>{auth.first_name}</strong>
									{/* Hello, <strong>Name</strong> */}
								</span>
								<Button variant='outline-dark' size='sm' onClick={handleLogout}>
									Logout
								</Button>
							</>
						) : (
							<>
								<Button
									variant='outline-info'
									size='sm'
									className='me-2'
									onClick={() => navigate('/login')}
								>
									Login
								</Button>
								<Button
									variant='info'
									size='sm'
									onClick={() => navigate('/register')}
								>
									Register
								</Button>
							</>
						)}
					</Nav>
				</Navbar.Collapse>
			</Container>
		</Navbar>
	)
}
export default Header
