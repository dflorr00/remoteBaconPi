import React, { useState, useEffect } from "react";
import {
  ListItemButton,
  ListItem,
  Typography,
  Grid,
  Box,
  List,
  CssBaseline,
  AppBar,
  Toolbar,
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
} from "@mui/material";
import MenuLateral from "../../components/MenuLateral";
import { useUsers } from "../../hooks/useUsers";
import { useViews } from "../../hooks/useViews";

export default function Views() {
  const drawerWidth = 200;
  const { users, getUsers } = useUsers();
  const { views,addView, getViews, deleteView } = useViews();
  const [userSelec, setUserSelec] = useState(null);
  const [visible, setVisible] = useState("hidden");

  const eliminar = (deviceId, ownerId) => {
    deleteView(deviceId, ownerId);
    getViews();
  };

  const addViewClick = (userId,deviceId) => {
    addView(userId,deviceId);
  }

  const handleClick = (id) => {
    setVisible("visible");
    setUserSelec(id);
  };

  // Al cargar la pagina
  useEffect(() => {
    getViews();
    getUsers();
  }, []);

  const filteredViews = views.filter(
    (view) => userSelec === 1 || view.UserID === userSelec
  );

  return (
    <Box sx={{ display: "flex" }}>
      <CssBaseline />
      <AppBar
        position="fixed"
        sx={{ width: `calc(100% - ${drawerWidth}px)`, ml: `${drawerWidth}px` }}
      >
        <Toolbar>
          <Typography variant="h6" noWrap component="div">
            VISTAS DE USUARIOS
          </Typography>
        </Toolbar>
      </AppBar>
      <MenuLateral />
      <Box sx={{ flexGrow: 1, bgcolor: "background.default", p: 3 }}>
        <Toolbar />
        <Grid container spacing={2}>
          <Grid item xs={12} sm={6}>
            <Typography variant="h6" align="center">
              USUARIOS
            </Typography>
            <List>
              {users.map((user, index) => (
                <ListItemButton
                  key={index}
                  onClick={() => handleClick(user.UserID)}
                  sx={{
                    bgcolor: "white",
                    borderRadius: 4,
                    boxShadow: 1,
                    p: 2,
                    mb: 2,
                  }}
                >
                  <Grid container spacing={1}>
                    <Grid item xs={12} sm={4}>
                      <Typography variant="subtitle1" color="text.secondary">
                        Id:
                      </Typography>
                      <Typography>{user.UserID}</Typography>
                    </Grid>
                    <Grid item xs={12} sm={4}>
                      <Typography variant="subtitle1" color="text.secondary">
                        Usuario:
                      </Typography>
                      <Typography>{user.UserName}</Typography>
                    </Grid>
                  </Grid>
                </ListItemButton>
              ))}
            </List>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Typography
              variant="h6"
              align="center"
              sx={{ visibility: visible }}
            >
              VISTAS DEL USUARIO {userSelec}
            </Typography>
            <List>
              {filteredViews
                .sort((a, b) => a.UserID - b.UserID)
                .reduce((acc, view) => {
                  const index = acc.findIndex(
                    (item) => item.DeviceID === view.DeviceID
                  );
                  if (index !== -1) {
                    acc[index].UserID.push(view.UserID);
                  } else {
                    acc.push({ ...view, UserID: [view.UserID] });
                  }
                  return acc;
                }, [])
                .map((view, index) => (
                  <ListItem
                    key={index}
                    sx={{
                      borderRadius: 4,
                      boxShadow: 1,
                      p: 2,
                      mb: 2,
                    }}
                  >
                    <Grid container spacing={1}>
                      <Grid item xs={12} sm={4}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Id del dispositivo: <b>{view.DeviceID}</b>
                        </Typography>
                      </Grid>
                      <Grid item xs={12} sm={4}>
                        <Typography variant="subtitle1" color="text.secondary">
                          Visible por:{" "}
                          <b>{view.UserID.sort((a, b) => a - b).join(", ")}</b>
                        </Typography>
                      </Grid>
                      <Grid item xs={12} sm={4}>
                        {userSelec !== 1 && "Visible:"}
                        {userSelec !== 1 && (
                          <Button
                            onClick={(e) => {
                              e.stopPropagation();
                              eliminar(view.DeviceID, userSelec);
                            }}
                          >
                            Eliminar
                          </Button>
                        )}
                        {userSelec === 1 && "Añadir: "}
                        {userSelec === 1 &&(
                            <FormControl>
                              <Select
                                onChange={(e) => {
                                  addViewClick(e.target.value,view.DeviceID);
                                }}
                              >
                                {users.filter((user) => !view.UserID.includes(user.UserID)).map((user) => (
                                  <MenuItem
                                    key={user.UserID}
                                    value={user.UserID}
                                  >
                                    {user.UserID} - {user.UserName}
                                  </MenuItem>
                                ))}
                              </Select>
                            </FormControl>
                          )}
                      </Grid>
                    </Grid>
                  </ListItem>
                ))}
            </List>
          </Grid>
        </Grid>
      </Box>
    </Box>
  );
}
