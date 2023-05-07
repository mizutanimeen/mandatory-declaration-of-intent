const baseURL = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "");
export const getRoomURL = (id: string) => { return baseURL + '/rooms/' + id };
export const postGestUserURL = () => { return baseURL + '/rooms/members/gest' };
export const getAllGestUserURL = (id: string) => { return baseURL + '/rooms/' + id + '/members/gest' };
export const postRoomURL = () => { return baseURL + '/rooms' };
export const getRoomPasswordCheckURL = (id: string) => { return baseURL + '/rooms/' + id + '/check' };

