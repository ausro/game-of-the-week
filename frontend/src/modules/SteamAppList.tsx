import axios from "axios"
import { useEffect, useRef, useState } from "react"
import { SteamAppDetail } from "../types/SteamApp"
import { API_URL } from "../util/urls"
import SteamApp from "./SteamApp"
import convDateToNum from "../util/convDateToNum"

interface SteamAppListProps {
    recommended: boolean
}

export default function SteamAppList({recommended}: SteamAppListProps) {
    const [appList, setAppList] = useState<SteamAppDetail[]>()
    const initialized = useRef(false)

    useEffect(() => {
        const getData = async(target: string) => {
            let list: SteamAppDetail[] = []
            await axios.get(API_URL + target).then(({data}) => {
                Object.keys(data).forEach(k => {
                    list.push(data[k])
                })

                list.sort((a, b) => convDateToNum(b.release_date) - convDateToNum(a.release_date))
                setAppList(list)
            })
        }

        function queryData() {
            if (!initialized.current) {
                if (recommended) {
                    getData("/recommended")
                } else {
                    getData("/")
                }
                initialized.current = true
            }
        }

        queryData()

    }, [])

    return (
        <div>
            <div className="container">
                <div className="grid">
                    {appList?.map((app) => <SteamApp app={app}></SteamApp>)}
                </div>
            </div>
        </div>
    )
}