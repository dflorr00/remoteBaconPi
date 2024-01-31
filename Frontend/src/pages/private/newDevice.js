import React,{useState} from "react";
import {Typography } from "@mui/material";
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Grid from '@mui/material/Grid';
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Checkbox from '@mui/material/Checkbox';
import FormControlLabel from '@mui/material/FormControlLabel';
import MenuLateral from "../../components/MenuLateral";
import { useDevices } from "../../hooks/useDevices";

export default function NewDevice() {
  const drawerWidth = 200;
  const [active, setActive] = useState(false);
  const {addDevice} = useDevices();

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);
    addDevice(data);
  };

  return (
    <Box sx={{ display: 'flex' }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
      >
        <Toolbar>
          <Typography variant="h6" noWrap component="div">
            AÑADIR DISPOSITIVO
          </Typography>
        </Toolbar>
      </AppBar>
      <MenuLateral/>
      <Box
        component="form"
        sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3 }}
        onSubmit={handleSubmit}
      >
        <Toolbar />
        <Grid container spacing={1} >
          <Grid item xs={5}><TextField margin="normal" required fullWidth label="Nombre del dispositivo" name="devicename" /></Grid>
          <Grid item xs={8}><TextField margin="normal" required fullWidth label="Dirección ip" name="ip" /></Grid>
          <Grid item xs={3}><TextField margin="normal" required fullWidth label="Puerto" name="port" /></Grid>
          <Grid item xs={12}><Typography><b>Servicios</b></Typography></Grid>
          <Grid item xs={12} sx={{ padding: "0px" }}>
            <FormControlLabel control={<Checkbox onChange={()=>setActive(!active)}/>} name='http' label="http" />
            <FormControlLabel control={<Checkbox onChange={()=>setActive(!active)}/>} name='printer' label="printer" />
          </Grid>
        </Grid>
        <Button type="submit" fullWidth variant="contained" disabled={!active} sx={{ mt: 2 }}> Añadir </Button>
      </Box>
    </Box>
  );
}