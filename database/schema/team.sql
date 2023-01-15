CREATE TABLE teams
(
    team_id INT PRIMARY KEY NOT NULL,
    key VARCHAR(3) NOT NULL,
    active BOOLEAN NOT NULL,
    city TEXT,
    name TEXT,
    league_id INT,
    stadium_id INT,
    conference TEXT,
    division TEXT,
    primary_color VARCHAR(6),
    secondary_color VARCHAR(6),
    tertiary_color VARCHAR(6),
    quaternary_color VARCHAR(6),
    wikipedia_logo_url TEXT,
    wikipedia_word_mark_url TEXT,
    global_team_id INT,
    nba_dot_com_team_id INT
);