import { useState } from "react";
import { useRouter } from "next/router";
const axios = require('axios')


export default function Create() {
  const router = useRouter();
  const [user, setUser] = useState({
    nombre: "",
    email: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUser({ ...user, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/users', user, {
        headers: {
          'Content-Type': 'application/json',
        }
      });
      router.push("/users");
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div>
      <h1>Create New User</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Name</label>
          <input type="text" name="nombre" value={user.nombre} onChange={handleChange} required />
        </div>
        <div>
          <label>Email</label>
          <input type="email" name="email" value={user.email} onChange={handleChange} required />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}
