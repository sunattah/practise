const API_URL = 'http://localhost:8080/api';

// Tab switching
function showTab(tabName) {
    const tabs = document.querySelectorAll('.tab-content');
    const btns = document.querySelectorAll('.tab-btn');
    
    tabs.forEach(tab => tab.classList.remove('active'));
    btns.forEach(btn => btn.classList.remove('active'));
    
    document.getElementById(tabName).classList.add('active');
    event.target.classList.add('active');
    
    if (tabName === 'all-students') {
        loadStudents();
    }
}

// Load all students
async function loadStudents() {
    try {
        const response = await fetch(`${API_URL}/students`);
        const students = await response.json();
        
        const container = document.getElementById('students-list');
        
        if (students.length === 0) {
            container.innerHTML = '<p>No students found. Add some students to get started!</p>';
            return;
        }
        
        container.innerHTML = students.map(student => `
            <div class="student-card">
                <h3>👤 ${student.name}</h3>
                <div class="student-info">
                    <div class="info-item"><strong>ID:</strong> ${student.id}</div>
                    <div class="info-item"><strong>Age:</strong> ${student.age}</div>
                    <div class="info-item"><strong>Grade:</strong> ${student.grade}</div>
                    <div class="info-item"><strong>Email:</strong> ${student.email}</div>
                    <div class="info-item"><strong>Phone:</strong> ${student.phone}</div>
                    <div class="info-item"><strong>Attendance:</strong> ${student.attendance}%</div>
                </div>
                <div class="action-buttons">
                    <button class="btn btn-success" onclick="viewDashboard(${student.id})">View Dashboard</button>
                    <button class="btn btn-danger" onclick="deleteStudent(${student.id})">Delete</button>
                </div>
            </div>
        `).join('');
    } catch (error) {
        console.error('Error loading students:', error);
    }
}

// Add student
document.getElementById('add-student-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const student = {
        name: document.getElementById('name').value,
        age: parseInt(document.getElementById('age').value),
        grade: document.getElementById('grade').value,
        email: document.getElementById('email').value,
        phone: document.getElementById('phone').value
    };
    
    try {
        const response = await fetch(`${API_URL}/add`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(student)
        });
        
        if (response.ok) {
            alert('✅ Student added successfully!');
            e.target.reset();
            loadStudents();
        }
    } catch (error) {
        console.error('Error adding student:', error);
        alert('❌ Error adding student');
    }
});

// Delete student
async function deleteStudent(id) {
    if (!confirm('Are you sure you want to delete this student?')) return;
    
    try {
        const response = await fetch(`${API_URL}/delete?id=${id}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            alert('🗑️ Student deleted successfully!');
            loadStudents();
        }
    } catch (error) {
        console.error('Error deleting student:', error);
    }
}

// View dashboard
function viewDashboard(id) {
    document.getElementById('dashboard-id').value = id;
    showTab('dashboard');
    document.querySelectorAll('.tab-btn')[2].classList.add('active');
    loadDashboard();
}

// Load dashboard
async function loadDashboard() {
    const id = document.getElementById('dashboard-id').value;
    
    if (!id) {
        alert('Please enter a student ID');
        return;
    }
    
    try {
        const response = await fetch(`${API_URL}/student?id=${id}`);
        
        if (!response.ok) {
            alert('Student not found');
            return;
        }
        
        const student = await response.json();
        
        const attendanceStatus = student.attendance >= 90 ? 'Excellent' : 
                                 student.attendance >= 75 ? 'Good' : 'Needs Improvement';
        
        const coursesHTML = student.courses && student.courses.length > 0 
            ? `<ul class="course-list">${student.courses.map(c => `<li>📚 ${c}</li>`).join('')}</ul>`
            : '<p>No courses enrolled yet.</p>';
        
        document.getElementById('dashboard-content').innerHTML = `
            <div class="dashboard-box">
                <div class="dashboard-section">
                    <h3>👤 Personal Information</h3>
                    <p><strong>Student ID:</strong> ${student.id}</p>
                    <p><strong>Name:</strong> ${student.name}</p>
                    <p><strong>Age:</strong> ${student.age} years old</p>
                    <p><strong>Email:</strong> ${student.email}</p>
                    <p><strong>Phone:</strong> ${student.phone}</p>
                    <p><strong>Enrolled Date:</strong> ${student.enrolledDate}</p>
                </div>
                
                <div class="dashboard-section">
                    <h3>📖 Academic Information</h3>
                    <p><strong>Current Grade:</strong> ${student.grade}</p>
                    <p><strong>Attendance:</strong> ${student.attendance}%</p>
                </div>
                
                <div class="dashboard-section">
                    <h3>📚 Enrolled Courses</h3>
                    ${coursesHTML}
                    <button class="btn btn-success" onclick="addCoursePrompt(${student.id})">Add Course</button>
                </div>
                
                <div class="dashboard-section">
                    <h3>📊 Quick Stats</h3>
                    <div class="stats-grid">
                        <div class="stat-box">
                            <div class="number">${student.courses ? student.courses.length : 0}</div>
                            <div>Total Courses</div>
                        </div>
                        <div class="stat-box">
                            <div class="number">${student.attendance}%</div>
                            <div>Attendance</div>
                        </div>
                        <div class="stat-box">
                            <div>${attendanceStatus}</div>
                            <div>Status</div>
                        </div>
                    </div>
                </div>
                
                <button class="btn btn-success" onclick="updateAttendancePrompt(${student.id})">Update Attendance</button>
            </div>
        `;
    } catch (error) {
        console.error('Error loading dashboard:', error);
    }
}

// Add course prompt
async function addCoursePrompt(id) {
    const course = prompt('Enter course name:');
    if (!course) return;
    
    try {
        const response = await fetch(`${API_URL}/addCourse`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id, course })
        });
        
        if (response.ok) {
            alert('✅ Course added successfully!');
            loadDashboard();
        }
    } catch (error) {
        console.error('Error adding course:', error);
    }
}

// Update attendance prompt
async function updateAttendancePrompt(id) {
    const attendance = prompt('Enter attendance percentage (0-100):');
    if (!attendance) return;
    
    try {
        const response = await fetch(`${API_URL}/updateAttendance`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id, attendance: parseFloat(attendance) })
        });
        
        if (response.ok) {
            alert('✅ Attendance updated successfully!');
            loadDashboard();
        }
    } catch (error) {
        console.error('Error updating attendance:', error);
    }
}

// Load students on page load
loadStudents();

