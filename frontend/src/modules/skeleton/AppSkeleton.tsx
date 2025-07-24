import Skeleton from '@mui/material/Skeleton'

const AppSkeleton = () => (
    <div className="container">
        <div className="grid">
            <a>
                <div className="app-panel" style={{display: 'grid', justifySelf: 'center', justifyItems: 'center'}}>
                    <Skeleton className="app-img" variant="rounded"  width={460} height={215}><img></img></Skeleton>
                    <Skeleton className="app-name" variant="text" width={200} height={50} />
                    <Skeleton className="app-desc" variant="text" width={300} height={30}/>
                    <Skeleton className="app-price" variant="text" width={100} height={30}/>
                    <Skeleton className="rel-date" variant="text" width={120} height={30}/>
                    <div className="genres" style={{ display: 'grid', width: "100%"}}>
                        {Array.from({ length: 3 }).map((_, i) => (
                            <Skeleton key={i} className="genre" variant="rounded" style={{marginLeft: 50, marginRight: 50, marginTop: 15}} width={60} height={28} />
                        ))}
                    </div>
                </div>
            </a>
        </div>
    </div>
)

export default AppSkeleton