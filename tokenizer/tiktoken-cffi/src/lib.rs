use std::ffi::CStr;

// https://doc.rust-lang.org/book/ch19-01-unsafe-rust.html#using-extern-functions-to-call-external-code
// https://web.mit.edu/rust-lang_v1.25/arch/amd64_ubuntu1404/share/doc/rust/html/book/first-edition/ffi.html

#[no_mangle]
pub extern "C" fn count_tokens(model: *const libc::c_char, prompt: *const libc::c_char) -> libc::c_uint {
    let model = unsafe { CStr::from_ptr(model).to_str().unwrap() };
    let prompt = unsafe { CStr::from_ptr(prompt).to_str().unwrap() };
    let bpe = tiktoken_rs::get_bpe_from_model(model).unwrap();
    let count = bpe.encode_with_special_tokens(prompt).len();
    count as libc::c_uint
}

#[no_mangle]
pub extern "C" fn get_context_size(model: *const libc::c_char) -> libc::c_uint {
    let model = unsafe { CStr::from_ptr(model).to_str().unwrap() };
    let size = tiktoken_rs::model::get_context_size(model);
    size as libc::c_uint
}
