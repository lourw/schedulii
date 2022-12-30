import '../App.css'
import { AppContextInterface, useAuth } from '../context/AuthContext'

const LandingPage = () => {
  const appContext: AppContextInterface | null = useAuth()
  
  return (
    <div className='wrapper'>
        <div className='header'></div>
        <div className='title'>
            <span>Schedulii</span>
            <button className='loginButton' onClick={appContext?.onLogin}>Login with Google</button>
        </div>

        <p>Token: { appContext?.token }</p>

        <div className='footer'></div>
    </div>
  )
}

export default LandingPage
