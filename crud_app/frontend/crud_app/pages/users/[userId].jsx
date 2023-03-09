import { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
const axios = require('axios')


const Update = ({single_user}) => {
  const [user, setUser] = useState({ nombre: single_user.nombre , email: single_user.email });
  const router = useRouter();
  const { userId } = router.query;


  const handleSubmit = async (event) => {
    event.preventDefault();

    await axios.put(`http://localhost:8080/users/${userId}`, user, {
      headers: {
        'Content-Type': 'application/json',
      },
    });

    router.push('/users');
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setUser({ ...user, [name]: value });
  };

  return (
    <div>
      <h1>Edit User: {single_user.nombre}</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            name="nombre"
            value={user.nombre}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={user.email}
            onChange={handleChange}
          />
        </div>
        <button>Save</button>
      </form>
    </div>
  );
};

export default Update;

export async function getServerSideProps(context) {
  const { userId } = context.query;
  const response = await axios.get(`http://localhost:8080/users/${userId}`);
  const single_user = await response.data;
  return {
    props: { single_user }, // will be passed to the page component as props
  }
}