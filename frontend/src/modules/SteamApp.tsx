import { STEAM_STORE } from "../util/urls"

interface SteamAppProps {
    app?: any
}

export default function SteamApp({ app }: SteamAppProps) {
    return (
        <a href={STEAM_STORE + app.id} target="_blank" rel="noopener noreferrer">
            <div className="app-panel">
                <img className="app-img" src={app.header_image} alt={app.name} />
                <p className="app-name">{app.name}</p>
                <p className="app-desc">{app.short_description}</p>
                <p className="app-price"><em>{app.price}</em></p>
                <p className="rel-date">Released: {app.release_date}</p>
                <div className="genres">
                    {app.genres.map((genre: string) =>
                        <div className="genre" key={genre}><p>{genre}</p></div>
                    )}
                </div>
            </div>
        </a>
    )
}