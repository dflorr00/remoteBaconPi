import { useState } from "react";
import axios from "axios";

export const useUsers = () =>{
    const [users, setUsers] = useState([]);
    
    const addUser = async (data) => {
      axios.post("http://127.0.0.1:8080/users/", {
        UserName: data.get('username'),
        Password: data.get('password'),
        Email: data.get('email')
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      }).then(response => {
        alert(response.data);
      }).catch(error => {
        alert("Error en la solicitud:", error);
      });
    }

    const getUsers = async () => {
        try {
        const response = await axios.get("http://127.0.0.1:8080/users/", {
            withCredentials: true,
        });
        console.log("Informacion de la respuesta", response.data);
        if (response.data != null) {
            setUsers(response.data);
        }
        console.log(users);
        } catch (error) {
        console.log(error);
        }
    };

    return {
        users,
        addUser,
        getUsers
    };

}  