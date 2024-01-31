import * as React from "react";
import PrintIcon from '@mui/icons-material/Print';
import LogoutIcon from '@mui/icons-material/Logout';
import Drawer from '@mui/material/Drawer';
import Divider from '@mui/material/Divider';
import ListItem from '@mui/material/ListItem';
import { ListItemButton} from "@mui/material";
import List from '@mui/material/List';
import AddIcon from '@mui/icons-material/Add';
import RemoveRedEyeIcon from '@mui/icons-material/RemoveRedEye';
import Toolbar from '@mui/material/Toolbar';
import { LogOut} from "../hooks/useServer";
import { useNavigate } from "react-router-dom";


export default function MenuLateral() {
  const drawerWidth = 200;
  const navigate = useNavigate();
  return (
    <div>
    <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
            boxSizing: 'border-box',
          },
        }}
        variant="permanent"
        anchor="left"
      >
        <Toolbar />
        <Divider />
        <List>
          <ListItem disablePadding>
            <ListItemButton href="/home">
              <PrintIcon sx={{mr:2}}/>Dispositivos
            </ListItemButton>
          </ListItem>
          <ListItem disablePadding>
            <ListItemButton href="/newDevice">
              <AddIcon sx={{mr:2}}/>Añadir dispositivo
            </ListItemButton>
          </ListItem>
          <ListItem disablePadding>
            <ListItemButton href="/views">
              <RemoveRedEyeIcon sx={{mr:2}}/>Vistas
            </ListItemButton>
          </ListItem>
        </List>
        <Divider />
        <List>
          <ListItem disablePadding>
              <ListItemButton onClick={()=>LogOut(navigate)}>
                  <LogoutIcon sx={{mr:2}}/>Salir
              </ListItemButton>
            </ListItem>
        </List>
      </Drawer>
    </div>
  );
}