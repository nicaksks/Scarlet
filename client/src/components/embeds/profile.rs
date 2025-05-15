use serenity::all::{CreateEmbed, CreateEmbedAuthor, CreateEmbedFooter};

use crate::constants;

use models::most_played_character::MostPlayedCharacter;
use models::profile::{Profile, User};
use scarlet::models;

pub struct EmbedProfile;

impl EmbedProfile {
    pub fn new(profile: User, character: MostPlayedCharacter) -> CreateEmbed {
        let Profile {
            nick_name,
            league,
            experiement_memory,
            labyrinth_point,
            tournament_point,
            watch_word,
            ..
        } = profile.user;

        let character_name =
            constants::get_character_name(character.character_play_count.most_played_character);

        CreateEmbed::new()
            .author(CreateEmbedAuthor::new(nick_name))
            .image(format!(
                "{}/skin/{:0>3}.png",
                constants::CDN,
                character.character_play_count.most_played_character
            ))
            .thumbnail("")
            .description(format!(
                "
                League Points: {}\n
                Experiement Memory: {}\n
                Labyrinth Point: {}\n
                Tournament Point: {}\n",
                constants::league_name(league),
                experiement_memory,
                labyrinth_point,
                tournament_point
            ))
            .fields(
                [
                    (
                        "Most Played Character".to_string(),
                        format!("{}", character_name),
                        true,
                    ),
                    (
                        "Wins".to_string(),
                        format!("{}", character.character_play_count.win_count),
                        true,
                    ),
                ]
                .into_iter(),
            )
            .color(constants::COLOR)
            .footer(CreateEmbedFooter::new(watch_word))
    }
}
