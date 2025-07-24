export interface SteamAppDetail {
    id: number;
    name: string;
    price: string;
    genres: string[];
    release_date: string;
    short_description: string;
    header_image: string;
    promoted: boolean;
}

export interface SteamAppResponse {
    [key: string]: SteamAppDetail;
}