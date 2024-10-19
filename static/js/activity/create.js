var __async = (__this, __arguments, generator) => {
  return new Promise((resolve, reject) => {
    var fulfilled = (value) => {
      try {
        step(generator.next(value));
      } catch (e) {
        reject(e);
      }
    };
    var rejected = (value) => {
      try {
        step(generator.throw(value));
      } catch (e) {
        reject(e);
      }
    };
    var step = (x) => x.done ? resolve(x.value) : Promise.resolve(x.value).then(fulfilled, rejected);
    step((generator = generator.apply(__this, __arguments)).next());
  });
};

// assets/js/config.js
var inputErrorClass = "has-error";
var successBgClass = "alert-success";
var errorBgClass = "alert-error";

// assets/js/components/notification.js
function showNotification(message, type) {
  const notification = document.getElementById("notification");
  const header = document.getElementById("notification-header");
  const body = document.getElementById("notification-message");
  if (notification && header && body) {
    if (type === "success") {
      header.textContent = "Action Completed";
      notification.classList.remove(errorBgClass);
      notification.classList.add(successBgClass);
    } else {
      header.textContent = "Action Failed";
      notification.classList.remove(successBgClass);
      notification.classList.add(errorBgClass);
    }
    body.textContent = message;
    notification.style.display = "block";
  }
}

// assets/js/form.js
function submitForm(form, cb) {
  return __async(this, null, function* () {
    clearFormErrors(form);
    const formData = new FormData(form);
    const actionUrl = form.getAttribute("action");
    const methodInput = form.querySelector('input[name="_method"]');
    let method = "POST";
    if (methodInput)
      method = methodInput.value.toUpperCase();
    const payload = {};
    formData.forEach((value, key) => {
      if (key.endsWith("_id")) {
        payload[key] = Number(value);
      } else {
        payload[key] = value;
      }
    });
    try {
      const response = yield fetch(actionUrl, {
        method,
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });
      const { errors, message, data } = yield response.json();
      if (!response.ok) {
        if (errors)
          displayFormErrors(form, errors);
        showNotification(message, "error");
      } else {
        showNotification(message, "success");
        if (method !== "PUT")
          form.reset();
        cb(data);
      }
    } catch (error) {
      showNotification("An error occurred. Please try again.", "error");
    }
  });
}
function displayFormErrors(form, errors) {
  errors.forEach(({ field, error }) => {
    const input = form.querySelector(`[name="${field}"]`);
    if (input) {
      const helpText = input.nextElementSibling;
      input.classList.add(inputErrorClass);
      if (helpText)
        helpText.textContent = error;
    }
  });
}
function clearFormErrors(form) {
  form.querySelectorAll("." + inputErrorClass).forEach((input) => {
    input.classList.remove(inputErrorClass);
    const nextEl = input.nextElementSibling;
    if (nextEl)
      nextEl.textContent = "";
  });
}

// assets/js/components/dialog.js
function mountDialogForSelect(dialogId, select) {
  const dialog = document.getElementById(dialogId);
  const dialogClose = dialog.querySelector("#dialog-close");
  dialogClose == null ? void 0 : dialogClose.addEventListener("click", () => dialog == null ? void 0 : dialog.close());
  window.addEventListener("click", (event) => {
    if (event.target === dialog)
      dialog.close();
  });
  select == null ? void 0 : select.addEventListener("change", showDialog);
  select == null ? void 0 : select.addEventListener("click", showDialog);
  function showDialog() {
    const optionValue = select.options[select.selectedIndex].value;
    if (optionValue === "add") {
      dialog == null ? void 0 : dialog.showModal();
    }
  }
  return dialog;
}

// assets/js/components/select.js
function updateSelect(select, detail) {
  const options = Array.from(select.options).slice(1);
  const newOption = document.createElement("option");
  newOption.value = detail.id.toString();
  newOption.text = detail.name;
  newOption.selected = true;
  options.push(newOption);
  options.sort((a, b) => a.text.localeCompare(b.text));
  const firstOption = select.options[0];
  select.innerHTML = "";
  select.add(firstOption);
  options.forEach((option) => select.add(option));
}

// assets/js/host.js
var createHostForm = document.getElementById("create-host-form");
var hostSelect = document.getElementById("host_id");
var hostDialog = mountDialogForSelect("create-host-dialog", hostSelect);
function watchHost() {
  hostSelect == null ? void 0 : hostSelect.addEventListener(
    "HostCreated",
    /** @param {import('./typedefs.js').MyCustomEventInit} event */
    function(event) {
      updateSelect(this, event.detail);
    }
  );
}
function handleHostForm() {
  createHostForm == null ? void 0 : createHostForm.addEventListener("submit", function(event) {
    event.preventDefault();
    submitForm(this, (data) => {
      hostSelect == null ? void 0 : hostSelect.dispatchEvent(
        new CustomEvent("HostCreated", { detail: data })
      );
      hostDialog.close();
    });
  });
}

// assets/js/venue.js
var createVenueForm = document.getElementById("create-venue-form");
var venueSelect = document.getElementById("venue_id");
function handleVenueForm() {
  createVenueForm == null ? void 0 : createVenueForm.addEventListener("submit", function(event) {
    event.preventDefault();
    submitForm(this, (data) => {
      venueSelect == null ? void 0 : venueSelect.dispatchEvent(
        new CustomEvent("VenueCreated", { detail: data })
      );
      venueDialog.close();
    });
  });
}
var venueDialog = mountDialogForSelect("create-venue-dialog", venueSelect);
function watchVenue() {
  venueSelect == null ? void 0 : venueSelect.addEventListener(
    "VenueCreated",
    /** @param {import('./typedefs.js').MyCustomEventInit} event */
    function(event) {
      updateSelect(this, event.detail);
    }
  );
}

// assets/js/activity/create.js
var createActivityForm = document.getElementById("create-activity-form");
createActivityForm == null ? void 0 : createActivityForm.addEventListener("submit", function(event) {
  event.preventDefault();
  submitForm(this, () => {
  });
});
handleVenueForm();
handleHostForm();
watchVenue();
watchHost();
//# sourceMappingURL=create.js.map
