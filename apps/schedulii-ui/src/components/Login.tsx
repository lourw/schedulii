import { useState } from "react";
import { useNavigate } from "react-router-dom";
function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = (event: React.FormEvent) => {
    event.preventDefault();
    console.log(email);
    if (email.trim().length > 0 && password.trim().length > 0) {
      navigate("/home");
    }
  };

  return (
    <div>
      {" "}
      <h2 className="text-2xl font-bold mb-3">Sign In</h2>
      <form className="space-y-2 flex flex-col" onSubmit={handleLogin}>
        <label className="text-slate-900">Email:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="email"
          placeholder="you@example.com"
          value={email}
          onChange={(event) => {
            setEmail(event.target.value);
          }}
        ></input>
        <label className="text-slate-900">Password:</label>
        <input
          className="text-slate-500 rounded border px-1"
          type="password"
          value={password}
          onChange={(event) => {
            setPassword(event.target.value);
          }}
        ></input>
        <div className="m-10 py-5">
          <button
            className="bg-sky-400 font-medium text-slate-50 px-10 py-2 self-center rounded-full"
            type="submit"
          >
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
}

export default Login;
