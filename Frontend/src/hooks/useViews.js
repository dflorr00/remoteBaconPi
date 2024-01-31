import { useState } from "react";
import axios from "axios";

export const useViews= () =>{

  const [views, setViews] = useState([]);

  const addView = async (userId,deviceId) => {
    axios.post("http://127.0.0.1:8080/views/", {
      UserID: userId,
	    DeviceID: deviceId
    }, {
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => {
      getViews();
    }).catch(error => {
      alert("Error en la solicitud:", error);
    });
  }

  const getViews = async () => {
    try {
        const response = await axios.get(`http://127.0.0.1:8080/views/`, {
            withCredentials: true,
        });
        console.log("Informacion de la respuesta", response.data);
        setViews(response.data);
    } catch (error) {
        console.log(error);
    }
  }

  const deleteView = async (deviceId, ownerId) => {
    console.log("El usuario: ",ownerId," el device: ",deviceId);
    axios.delete('http://localhost:8080/views/', {
      data: {
        UserID: ownerId,
        DeviceID: deviceId
      }
    }).then(response => {
      getViews();
    }).catch(error => {
      alert("Error en la solicitud:", error);
    });
  }

  return {
    views,
    addView,
    getViews,
    deleteView
  };

}  