package config

import raspao "github.com/wawandco/raspao/base"

/*
 * DefineRoutes gets called by Raspao to define routes inside your raspao application,
 * please you need to have this funtion here for raspao to work, DO not remove it :).
 */

var SharedApplication *raspao.App

var DefineRoutes = func(r *raspao.Router) {
	/*
	 * Use the provider raspao's Router object methods to declare
	 * your routes.
	 * For example:
	 *
	 * r.Get("/", "HomeController.Index")
	 * r.Post("/{id}", "UserController.Create")
	 * r.Put("/{id}", "UserController.Update")
	 * r.Delete("/{id}", "UserController.Destroy")
	 *
	 * You can use any of the following methods:
	 *   r.Get      // GET
	 *   r.Post     // POST
	 *   r.Put      // PUT
	 *   r.Delete   // DELETE
	 *   r.Patch    // PATCH
	 *   r.Copy     // COPY
	 *   r.Head     // HEAD
	 *   r.Options  // OPTIONS
	 *   r.Link     // LINK
	 *   r.Unlink   // UNLINK
	 *   r.Purge    // PURGE
	 *   r.Lock     // LOCK
	 *   r.Unlock   // UNLOCK
	 *   r.Propfind // PROPFIND
	 *   r.View     // VIEW
	 *
	 * Or as an alternative:
	 *  r.Map("GET", "/{id}", "Controller.Show") // Is equal to r.Get("/{id}", "Controller.Show")
	 */
	r.Get("/", "NotesController.Initialize")
	r.Path("/notes", func() {
		r.Get("/", "NotesController.Index")
		r.Get("/create", "NotesController.Create")
		r.Get("/{id}", "NotesController.Find")
		r.Get("/delete/{id}", "NotesController.Delete")
	})

	r.Path("/v1", func() {
		r.Get("/", "RootController.Index")
		r.Get("/principal", "PrincipalController.Index")
	})
}
