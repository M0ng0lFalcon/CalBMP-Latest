const ssa = window.sessionStorage;

export default {
  getItem(key) {
    try {
      return JSON.parse(ssa.getItem(key));
    } catch (err) {
      return null;
    }
  },
  setItem(key, val) {
    ssa.setItem(key, JSON.stringify(val));
  },
  clear() {
    ssa.clear();
  },
  keys() {
    return ssa.keys();
  },
  removeItem(key) {
    ssa.removeItem(key);
  },
};
