# Portuguese legislative elections 2024 Datasheet

## Data Links
 - **[Votes CSV Preview](https://github.com/jtomaspm/legislativas2024/blob/main/data/votesPreview.csv)**
 - **[Votes CSV Complete](https://github.com/jtomaspm/legislativas2024/blob/main/data/votes.csv)**
 - **[Votes Excel Complete](https://github.com/jtomaspm/legislativas2024/blob/main/data/votes.xlsx)**
 - **[Locations CSV](https://github.com/jtomaspm/legislativas2024/blob/main/data/locations.csv)**

## Data Source

 - *[Location data](https://github.com/jtomaspm/legislativas2024/blob/main/data/locations.csv)* extracted from *[legislativas 2024 website](https://www.legislativas2024.mai.gov.pt/resultados/territorio-nacional?local=LOCAL-500000)* using this *[js tampermonkey script](https://github.com/jtomaspm/legislativas2024/blob/main/scripts/fetchLocations.js)*
 - Data sorted using this *[python script](https://github.com/jtomaspm/legislativas2024/blob/main/scripts/sortLocations.py)*
 - *[Vote data](https://github.com/jtomaspm/legislativas2024/blob/main/data/votesPreview.csv)* fetched by locationId from *[legislativas 2024 api](https://www.legislativas2024.mai.gov.pt/frontend/data/TerritoryResults?territoryKey=LOCAL-500000&electionId=AR)* using this *[golang program](https://github.com/jtomaspm/legislativas2024/blob/main/votes/main.go)*

## Component Diagram

![Component Diagram](https://github.com/jtomaspm/legislativas2024/blob/main/data/compDiagram.png)