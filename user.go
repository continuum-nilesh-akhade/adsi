package adsi

import (
	"github.com/continuum-nilesh-akhade/adsi/api"
	ole "github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
)

// User provides access to Active Directory users.
type User struct {
	object
	iface *api.IADsUser
}

// NewUser returns a user that manages the given COM interface.
func NewUser(iface *api.IADsUser) *User {
	comshim.Add(1)
	return &User{iface: iface, object: object{iface: &iface.IADs}}
}

func (u *User) closed() bool {
	return (u.iface == nil)
}

// Close will release resources consumed by the user. It should be
// called when the user is no longer needed.
func (u *User) Close() {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return
	}
	defer comshim.Done()
	u.iface.Release()
	u.object.iface = nil
	u.iface = nil
}

// FullName retrieves the FullName of the user.
func (u *User) FullName() (fullName string, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return "", ErrClosed
	}
	fullName, err = u.iface.FullName()
	return
}

// Description retrieves the description of the user.
func (u *User) Description() (description string, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return "", ErrClosed
	}
	description, err = u.iface.Description()
	return
}

// PasswordRequired retrieves the attribute of the user
func (u *User) PasswordRequired() (passReq bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	passReq, err = u.iface.PasswordRequired()
	return
}

// AccountDisabled retrieves the attribute of the user
func (u *User) AccountDisabled() (accDisabled bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	accDisabled, err = u.iface.AccountDisabled()
	return
}

// IsAccountLocked retrieves the attribute of the user
func (u *User) IsAccountLocked() (accLocked bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	accLocked, err = u.iface.IsAccountLocked()
	return
}

// RequireUniquePassword retrieves the attribute of the user
func (u *User) RequireUniquePassword() (reqUniqPass bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	reqUniqPass, err = u.iface.RequireUniquePassword()
	return
}

// PasswordMinimumLength retrieves the attribute of the user
func (u *User) PasswordMinimumLength() (reqUniqPass int64, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return 0, ErrClosed
	}
	reqUniqPass, err = u.iface.PasswordMinimumLength()
	return
}

// LastLogin retrieves the attribute of the user
//TODO: Convert int64 to valid Time and change return type to time.Time
func (u *User) LastLogin() (lastLogin int64, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return 0, ErrClosed
	}
	largeInt, err := u.iface.LastLogin()
	if err != nil {
		return
	}

	lastLogin, err = largeInt.Value()
	if err != nil {
		return
	}
	return
}

// PasswordExpirationDate retrieves the attribute of the user
//TODO: Convert int64 to valid Time and change return type to time.Time
func (u *User) PasswordExpirationDate() (passExp int64, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return 0, ErrClosed
	}
	passExp, err = u.iface.PasswordExpirationDate()
	if err != nil {
		return
	}
	return passExp, nil
}

// Get retrieves the property value of the user
//TODO: Add separate functions to convert ole.VARIANT to go types
func (u *User) Get(name string) (val *ole.VARIANT, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return nil, ErrClosed
	}
	val, err = u.iface.Get(name)
	if err != nil {
		return
	}
	return val, nil
}
