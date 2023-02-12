/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.17
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdFwStatsOk. */
//easyjson:json
type GetCharactersCharacterIdFwStatsOkList []GetCharactersCharacterIdFwStatsOk

/* 200 ok object */
//easyjson:json
type GetCharactersCharacterIdFwStatsOk struct {
	CurrentRank   int32                                        `json:"current_rank,omitempty"` /* The given character's current faction rank */
	EnlistedOn    time.Time                                    `json:"enlisted_on,omitempty"`  /* The enlistment date of the given character into faction warfare. Will not be included if character is not enlisted in faction warfare */
	FactionId     int32                                        `json:"faction_id,omitempty"`   /* The faction the given character is enlisted to fight for. Will not be included if character is not enlisted in faction warfare */
	HighestRank   int32                                        `json:"highest_rank,omitempty"` /* The given character's highest faction rank achieved */
	Kills         GetCharactersCharacterIdFwStatsKills         `json:"kills,omitempty"`
	VictoryPoints GetCharactersCharacterIdFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
