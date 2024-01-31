import * as React from 'react';
import Avatar from '@mui/material/Avatar';
import Button from '@mui/material/Button';
import TextField from '@mui/material/TextField';
import Box from '@mui/material/Box';
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { useNavigate } from 'react-router';
import theme from '../../theme';
import { LogIn } from '../../hooks/useServer';

export default function SignIn() {
  
  const navigate = useNavigate();

  const handleLogin = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    LogIn(data, navigate);
  };

  return (
    <Container component="main" maxWidth="xs">
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        {/*Logo y titulo*/}
        <Avatar sx={{ m: 1, bgcolor: theme.palette.primary.main }}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component="h1" variant="h5">
          Iniciar sesión
        </Typography>
        <div>
          {/*Inicio de sesión*/}
          <Box component="form" onSubmit={handleLogin} noValidate sx={{ mt: 1 }}>
            <TextField margin="normal" fullWidth label="Nombre de usuario" name="username" autoFocus />
            <TextField margin="normal" fullWidth label="Contraseña" name="password" type="password" autoComplete="current-password"/>
            <Button type="submit" fullWidth variant="contained" sx={{ mt: 2 }}>
              {' '}Iniciar sesión{' '}
            </Button>
          </Box>
          {/*Registro*/}
          <Box component="form" onSubmit={() => navigate('/signUp')}>
            <Button type="submit" fullWidth variant="contained" sx={{ mb: 2 }}>
              {' '}Registrarse{' '}
            </Button>
          </Box>
        </div>
      </Box>
      <div><a href='/control'>Ir a la página de control</a></div>
    </Container>
  );
}
