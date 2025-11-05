package controllers

import (
	ent "Friendy/entities"
	"fmt"
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// ApiControllerV1 handles version 1 of the Friend API endpoints.
type ApiControllerV1 struct {}

// ==========================
// PROTECTED CONTROLLER METHODS (Auth Required)
// ==========================

/* 
 CreateFriend creates a new friend record for the authenticated user.
 It binds the incoming JSON body into a map, attaches the current user's ID,
 creates a new record, and saves it to the "Friends" collection.
*/
func (r *ApiControllerV1) CreateFriend(e *core.RequestEvent) error {
	coll, err  := e.App.FindCollectionByNameOrId("Friends")
	if err != nil {
		return err
	}

	// Initialize map with user_id set to the authenticated user's ID
	friendMap := map[string]any{
		"user_id": e.Auth.Id,
	}

	// Merge request body fields into the map
	if err := e.BindBody(&friendMap); err != nil {
		return err
	}

	// Create a new record from the combined data
	record := core.NewRecord(coll)
	record.Load(friendMap)

	// Save the record into the database
	if err := e.App.Save(record); err != nil {
		return err
	}

	// Respond with the new record's ID
	return e.JSON(http.StatusOK, map[string]string{
		"friendId": record.Id,
	})

}

/* 
 ReadFriend retrieves a specific friend record belonging to the authenticated user.
 It validates ownership to ensure users can only read their own friend records.
*/ 
func (r *ApiControllerV1) ReadFriend(e *core.RequestEvent) error {
	var friend ent.Friend

	// Bind the incoming JSON to a Friend struct
	if err := e.BindBody(&friend); err != nil {
		return err
	}

	// Query the friend record by ID and ensure it belongs to the user
	err := e.App.DB().
	Select("*").
	From("Friends").
	Where(dbx.Like("id", friend.Id)).
	AndWhere(dbx.Like("user_id", e.Auth.Id)).
	One(&friend)

	if err != nil {
		return err
	}

	// Respond with the friend record
	return e.JSON(http.StatusOK, map[string]ent.Friend{
		"friend": friend,
	})

}

/* 
 UpdateFriend updates allowed fields of a friend record belonging to the authenticated user.
 It ensures ownership and prevents users from modifying other users' data.
*/ 
func (r *ApiControllerV1) UpdateFriend(e *core.RequestEvent) error {
	var friendUpdateMap map[string]any

	// Bind request body into a map
	if err := e.BindBody(&friendUpdateMap); err != nil {
		return err
	}

	// Ensure the record ID is provided
	friendId, ok := friendUpdateMap["id"]
	if !ok {
		return e.BadRequestError("Friend 'id' to update not provided", nil)
	}

	// Find the target record
	record, err := e.App.FindRecordById("Friends", friendId.(string))
	if err != nil {
		return err
	}

	// Verify that the current user owns this record
	if record.Get("user_id") != e.Auth.Id {
		return e.ForbiddenError("Can't access resource without owning it", nil)
	}

	// Define the fields that can be updated
	updatableFieldsMap := map[string]any{
		"fullname": nil,
		"tel": nil,
		"desc": nil,
		"first_met_on": nil,
		"met_place": nil,
		"tags": nil,
	}

	// Apply updates only to allowed fields
	for k, v := range friendUpdateMap {
		if _, ok := updatableFieldsMap[k]; ok {
			record.Set(k, v)
		}
	}

	// Save the updated record
	fmt.Println(record)
	if err := e.App.Save(record); err != nil {
		return err
	}

	// Respond with the updated friend data
	return e.JSON(http.StatusOK, map[string]ent.Friend{
		"friend": {
			Id: record.Id,
			FullName: record.GetString("fullname"),
			Tel: record.GetString("tel"),
			Desc: record.GetString("desc"),
			FirstMetOn: record.GetString("first_met_on"),
			MetPlace: record.GetString("met_place"),
			Tags: record.GetString("tags"),
			UserId: record.GetString("user_id"),
		},
	})

}

/* 
 DeleteFriend removes a friend record belonging to the authenticated user.
 It checks ownership before deleting to prevent unauthorized deletions.
*/ 
func (r *ApiControllerV1) DeleteFriend(e *core.RequestEvent) error {
	var friend ent.Friend

	// Bind request body into Friend struct
	if err := e.BindBody(&friend); err != nil {
		return err
	}

	// Retrieve the record by ID
	record, err := e.App.FindRecordById("Friends", friend.Id)
	if err != nil {
		return err
	}

	// Verify the authenticated user owns the record
	if record.Get("user_id") != e.Auth.Id {
		return e.ForbiddenError("Can't access resource without owning it", nil)
	}

	// Delete the record
	if err := e.App.Delete(record); err != nil {
		return err
	}

	// Respond with confirmation
	return e.JSON(http.StatusOK, map[string]string{
		"friendId": friend.Id,
	})

}

/* 
 AllFriends lists all friends belonging to the authenticated user.
 It limits results to 10 entries and returns a lightweight friend list.
*/ 
func (r *ApiControllerV1) AllFriends(e *core.RequestEvent) error {
	var friends []ent.Friend

	// Query friends belonging to the current user
	err := e.App.DB().
	Select("id", "fullname", "tel", "desc", "met_place").
	From("friends").
	Where(dbx.Like("user_id", e.Auth.Id)).
	Limit(10).
	All(&friends)

	if err != nil {
		return err 
	}

	// Respond with the friend list
	return e.JSON(http.StatusOK, map[string][]ent.Friend{
		"friends": friends,
	})

}

