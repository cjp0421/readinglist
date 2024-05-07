import { AppBar, Toolbar, Button, useMediaQuery, IconButton, Menu, MenuItem } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { useTheme } from '@mui/material/styles';
import { Link } from 'react-router-dom';
import { useState } from 'react';

const NavBar = () => {
    const theme = useTheme();
    const isMobile = useMediaQuery((theme.breakpoints.down('sm')));
    const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

    const handleMenu = (event) => {
        setAnchorEl(event.currentTarget);
    }

    const handleClose = () => {
        setAnchorEl(null)
    }

    return (
        <AppBar position="fixed">
            <Toolbar>
                {
                    isMobile ? (
                        <>
                            <IconButton
                                color="inherit"
                                aria-label="menu"
                                onClick={handleMenu}
                            >
                                <MenuIcon />
                            </IconButton>
                            <Menu
                                id="meun-appbar"
                                anchorEl={anchorEl}
                                anchorOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                keepMounted
                                transformOrigin={{
                                    vertical: 'top',
                                    horizontal: 'right',
                                }}
                                open={Boolean(anchorEl)}
                                onClose={handleClose}
                            >
                                <MenuItem onClick={handleClose} component={Link} to="/">Home</MenuItem>
                                <MenuItem onClick={handleClose} component={Link} to="/books/add">Add Book</MenuItem>
                            </Menu>
                        </>
                    ) : (
                        <>
                            <Button color="inherit" component={Link} to="/">Home</Button>
                            <Button color="inherit" component={Link} to="/books/add">Add Book</Button>
                        </>
                    )
                }

            </Toolbar>
        </AppBar>
    )
}

export default NavBar;