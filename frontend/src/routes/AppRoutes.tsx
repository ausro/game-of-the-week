import { lazy } from "react"
import {Route, Routes} from "react-router-dom"

const SteamAppList = lazy(() => import("../modules/SteamAppList"))
const About = lazy(() => import("../modules/About"))

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<SteamAppList recommended={false}/>}> </Route>
            <Route path="/recommended" element={<SteamAppList recommended={true}/>}></Route>
            <Route path="/about" element={<About />}></Route>
        </Routes>
    )
}

export default AppRoutes