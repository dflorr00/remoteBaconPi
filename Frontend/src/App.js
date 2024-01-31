import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Navigate, Routes, Route } from 'react-router-dom';
import SignUp from './pages/public/signUp';
import SignIn from './pages/public/signIn'
import Home from './pages/private/home';
import NewDevice from './pages/private/newDevice';
import Views from './pages/private/views';
import Device from './pages/private/device';
import { ThemeProvider } from '@mui/material/styles';
import theme from './theme';
import ProtectedRoute from './ProtectedRoute';
import Control from './pages/public/control';

function App() {

  return (
    <ThemeProvider theme={theme}>
      <Router>
        <Routes>
          <Route path="/" element={<Navigate to="/control" />} />
          <Route path="/control" element={<Control />} />
          <Route path="/signIn" element={<SignIn />} />
          <Route path="/signUp" element={<SignUp />} />
          <Route element={<ProtectedRoute />}>
            <Route path="/home" element={<Home />} />
            <Route path="/newDevice" element={<NewDevice />} />
            <Route path="/views" element={<Views />} />
            <Route path="/device/:id" element={<Device />} />
          </Route>
        </Routes>
      </Router>
    </ThemeProvider>
  );
}

export default App;
