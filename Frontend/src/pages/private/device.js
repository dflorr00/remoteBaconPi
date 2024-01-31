import React, { useState, useEffect } from "react";
import { Typography, Box, Grid, Paper, IconButton } from "@mui/material";
import axios from "axios";
import CssBaseline from '@mui/material/CssBaseline';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import MenuLateral from "../../components/MenuLateral";
import { useParams } from 'react-router-dom';
import DeleteIcon from '@mui/icons-material/Delete';
import { useDevices } from "../../hooks/useDevices";

export default function Device() {
  const { id } = useParams();
  const drawerWidth = 200;
  const {devices, getDevices} = useDevices();

  useEffect(() => {getDevices()}, []);

  const getSemafaroColor = (deviceStatus) => {
    switch (deviceStatus) {
      case 0:
          return "green";
      case 1:
          return "yellow";
      case 2:
          return "red";
      default: return "grey";    
    }
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
            ESTADO DE LA IMPRESORA {id}
          </Typography>
        </Toolbar>
      </AppBar>
      <MenuLateral />
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          bgcolor: 'background.default',
          p: 3,
          marginLeft: `${drawerWidth}px`,
        }}
      >
        <Toolbar />
        <Grid container spacing={2}>
          {devices.map((impresora) => {
            if (parseInt(impresora.DeviceID) === parseInt(id)) {
              return (
                <Grid item key={impresora.DeviceID} xs={12} sm={6} md={4} lg={3}>
                  <Paper elevation={3} sx={{ p: 2, borderRadius: 4, boxShadow: 1 }}>
                    <Grid container bgcolor="white" alignItems="center">
                      <Grid item xs={12}>
                        <Typography variant="subtitle1" color="text.secondary">
                          ID:
                        </Typography>
                        <Typography>{impresora.DeviceID}</Typography>
                      </Grid>
                      <Grid item xs={12}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Nombre:
                        </Typography>
                        <Typography>{impresora.DeviceName}</Typography>
                      </Grid>
                      <Grid item xs={12}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Dirección IP:
                        </Typography>
                        <Typography>{impresora.Ip}</Typography>
                      </Grid>
                      <Grid item xs={12}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Servicio:
                        </Typography>
                        <Typography>{impresora.Service}</Typography>
                      </Grid>
                      <Grid item xs={12}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Puerto:
                        </Typography>
                        <Typography>{impresora.Port}</Typography>
                      </Grid>
                      <Grid item xs={12} sx={{ display: 'flex', alignItems: 'center', justifyContent: 'flex-end' }}>
                        <div
                          style={{
                            width: 20,
                            height: 20,
                            borderRadius: '50%',
                            marginRight: 8,
                            backgroundColor: getSemafaroColor(impresora.Status),
                          }}
                        />
                      </Grid>
                    </Grid>
                  </Paper>
                </Grid>
              );
            }
            return null;
          })}
        </Grid>
      </Box>
    </Box>
  );
}
