namespace DeletePayment.model
{
    public class Payment
    {
        public int Id { get; set; }
        public decimal Amount { get; set; }
        public string Description { get; set; } = ""; // ✅ Inicializar la propiedad
        public DateTime Date { get; set; } = DateTime.Now; // ✅ Inicializar la propiedad
    }
}