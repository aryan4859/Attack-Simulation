<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Company Landing Page</title>
    <link rel="stylesheet" href="style.css">
</head>

<body>

    <!-- Navigation Bar -->
    <header>
        <div class="logo">
            <img src="logo.png" alt="Company Logo">
        </div>
        <nav>
            <ul>
                <li><a href="#home">Home</a></li>
                <li><a href="#services">Services</a></li>
                <li><a href="#about">About Us</a></li>
                <li><a href="#contact">Contact</a></li>
                <li><a href="#upload" class="upload-cv">Upload CV</a></li>
            </ul>
        </nav>
    </header>

    <!-- Hero Section -->
    <section id="home" class="hero">
        <div class="hero-content">
            <h1>Welcome to Our Company</h1>
            <p>We provide innovative solutions to help you succeed.</p>
            <a href="#contact" class="cta-button">Get in Touch</a>
        </div>
    </section>

    <!-- Services Section -->
    <section id="services" class="services">
        <h2>Our Services</h2>
        <div class="service-item">
            <h3>Consulting</h3>
            <p>We offer expert consulting services to help your business grow.</p>
        </div>
        <div class="service-item">
            <h3>Development</h3>
            <p>Our development team creates powerful applications and solutions.</p>
        </div>
        <div class="service-item">
            <h3>Support</h3>
            <p>We provide ongoing support to ensure the smooth running of your business.</p>
        </div>
    </section>

    <!-- About Section -->
    <section id="about" class="about">
        <h2>About Us</h2>
        <p>We are a passionate team dedicated to helping businesses succeed through technology and innovation.</p>
    </section>

    <!-- Contact Section -->
    <section id="contact" class="contact">
        <h2>Contact Us</h2>
        <p>Get in touch with us today and see how we can help your business grow.</p>
    </section>

    <!-- CV Upload Section -->
    <section id="upload" class="upload">
        <h2>Upload Your CV</h2>
        <form action="upload_cv.php" method="POST" enctype="multipart/form-data">
            <label for="cv">Select your CV (PDF or DOCX):</label>
            <input type="file" name="cv" id="cv" required>
            <button type="submit" name="submit">Upload CV</button>
        </form>
    </section>

    <!-- Footer -->
    <footer>
        <p>&copy; 2025 Our Company | <a href="#">Privacy Policy</a> | <a href="#">Terms of Service</a></p>
        <div class="social-media">
            <a href="#">Facebook</a> | <a href="#">LinkedIn</a> | <a href="#">Twitter</a>
        </div>
    </footer>

</body>

</html>
