import { lazy } from "react"
import {Route, Routes} from "react-router-dom"

const SteamAppList = lazy(() => import("../modules/SteamAppList"))

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<SteamAppList recommended={false}/>}> </Route>
            <Route path="/recommended" element={<SteamAppList recommended={true}/>}></Route>
        </Routes>
    )
}

export default AppRoutes