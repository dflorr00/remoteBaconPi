import React, {useEffect } from "react"
import {
  Card,
  CardContent,
  ListItem,
  Typography,
} from "@mui/material"
import Button from "@mui/material/Button"
import Box from "@mui/material/Box"
import List from "@mui/material/List"
import Grid from "@mui/material/Grid"
import Paper from "@mui/material/Paper"
import AddIcon from "@mui/icons-material/Add"
import DeleteIcon from "@mui/icons-material/Delete"
import CssBaseline from "@mui/material/CssBaseline"
import AppBar from "@mui/material/AppBar"
import Toolbar from "@mui/material/Toolbar"
import MenuLateral from "../../components/MenuLateral"
import { useNavigate } from "react-router"
import { useDevices } from "../../hooks/useDevices"

export default function Control() {
  const drawerWidth = 0
  const {devices, getDevices, deleteDevice} = useDevices()

  useEffect(() => {
    const intervalId = setInterval(() => {
        getDevices();
    }, 30000);
    return () => clearInterval(intervalId);
  }, []);
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
    <Box sx={{ display: "flex", flexDirection: "column"}}>
  <CardContent>
    <Card>
    <Typography variant="h4"  align="center" noWrap>
      ESTADO DE LAS IMPRESORAS
    </Typography>
    </Card>
  </CardContent>
      <Box component="main" sx={{ flexGrow: 1,p:3, bgcolor: "background.default"}}>
        <Grid container>
          <Grid item xs={10} sx={{ ml: 10, height: "65%"}}>
            <Paper square sx={{ bgcolor: "#0096D6" }}>
              <List sx={{ minWidth: "100%" }}>
                {devices.map((device, index) => (
                  <ListItem
                    key={index}
                    sx={{ minWidth: "100%" }}
                  >
                    <Grid
                      container
                      bgcolor="white"
                      alignItems="center"
                      sx={{ borderRadius: 4, boxShadow: 1, p: 2 }}
                    >
                      <Grid item xs={12} sm={6} md={4} lg={2}>
                        <Typography variant="subtitle1" color="text.secondary">
                          ID:
                        </Typography>
                        <Typography>{device.DeviceID}</Typography>
                      </Grid>
                      <Grid item xs={12} sm={6} md={4} lg={2}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Nombre:
                        </Typography>
                        <Typography sx={{ overflow: 'hidden', textOverflow: 'ellipsis' }}>{device.DeviceName}</Typography>
                      </Grid>
                      <Grid item xs={12} sm={6} md={4} lg={2}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Dirección IP:
                        </Typography>
                        <Typography>{device.Ip}</Typography>
                      </Grid>
                      <Grid item xs={12} sm={6} md={4} lg={2}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Servicio:
                        </Typography>
                        <Typography>{device.Service}</Typography>
                      </Grid>
                      <Grid item xs={12} sm={6} md={4} lg={2}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Puerto:
                        </Typography>
                        <Typography>{device.Port}</Typography>
                      </Grid>
                      <Grid
                        item
                        xs={12}
                        sm={6}
                        md={4}
                        lg={2}
                        sx={{ display: "flex", alignItems: "center", justifyContent: "flex-end" }}
                      >
                        <div
                          style={{
                            width: 20,
                            height: 20,
                            borderRadius: "50%",
                            marginRight: 8,
                            backgroundColor: getSemafaroColor(device.Status), // Función que devuelve el color correspondiente
                          }}
                        />
                        <Button
                          onClick={(e) => {
                            e.stopPropagation();
                            deleteDevice(device.DeviceID);
                          }}
                          variant="contained"
                          color="error"
                          sx={{ textTransform: "none", width: "100%" }}
                          startIcon={<DeleteIcon />}
                        >
                          Eliminar
                        </Button>
                      </Grid>
                      <Grid item xs={12}></Grid>
                      {/* Espacio en blanco para separar las tarjetas */}
                    </Grid>
                  </ListItem>
                ))}
              </List>
              <CardContent>
                <Card align="center">
                  <Button fullWidth href="signIn" variant="contained">Iniciar Sesión</Button>
                </Card>
              </CardContent>
            </Paper>
          </Grid>
        </Grid>
      </Box>
    </Box>
  );
}
